// goduino is miner for DUCO (https://duinocoin.com/) that support emulating your other devices to optimize hashrate

package main

import (
	// "fmt"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"strconv"
	"syscall"
	"time"
	// "time"
)

// usage: 	goduino mine [optional: -workers=how_many_devices] [optional: -rig_name=rig_name] [optional: -software_name=software_name] [optional: -debug true] -name=username -diff=difficulty (AVR for ????, LOW (for web miners and low-spec devices), MEDIUM (pc), NET (network diff - pc), EXTREME (officially not used anywhere, but you can use it to mine or gpus - howewer this program currently don't support mining on gpus))
//			or: goduino mine -config=./config.{json,yaml} # to mine with configuration file specified after --config

func cleanup(){
	log.Println("SIGINT has been detected, cleaning up and closing program")
	pprof.StopCPUProfile()
}

// handle SIGINT signal by capturing it and closing program 
func handleSIGINT(){
	c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        cleanup()
        os.Exit(1)
    }()
}

func main(){
	handleSIGINT()
	// Get args from flagS
	flag_workers := flag.Int("workers", 1, "how many devices you want to use")
	flag_rig_name := flag.String("rig_name", "goduino " + VERSION + " worker", "custom rig name")
	flag_software_name := flag.String("software_name", "goduino " + VERSION, "custom mining software name")
	flag_debug := flag.Bool("debug", false, "enable debug logs")
	flag_name := flag.String("name", "", "your username")
	flag_diff := flag.String("diff", "MEDIUM", "difficulty that you want to mine on")
	flag_config := flag.String("config", "", "enter path to your configuration file")
	flag_xxhash := flag.Bool("xxhash", false, "do you want to use xxhash instead of sha1")
	flag_max_hashrate := flag.String("max_hashrate", "", "do you want to limit hashrate to some value (supported options: LOW_NANO, LOW_ESP32, LOW_PI4; or: value in H/s eg. 720 will mean 720H/s)")
	flag_max_rejected := flag.Int("after how many rejected hashes you want to exit", 0, "default: disabled")


	flag.Parse()

	var workers int = *flag_workers - 1
	var rig_name string = *flag_rig_name
	var software_name string = *flag_software_name
	var debug bool = *flag_debug
	var username string = *flag_name
	var difficulty string = *flag_diff
	var config_path string = *flag_config
	var xxhash bool = *flag_xxhash
	var max_hashrate_str string = *flag_max_hashrate
	var max_rejected int = *flag_max_rejected
	
	// Set up logging
	if(!debug){log.SetOutput(ioutil.Discard)}
	
	log.Println("workers:", workers)
	log.Println("rig_name:", rig_name)
	log.Println("software_name:", software_name)
	log.Println("debug:", debug)
	log.Println("username:", username)
	log.Println("config_path:", config_path) 
	log.Println("max_hashrate_str:", max_hashrate_str)
	log.Println("VERSION:", VERSION)

	//read config //TODO ADD SUPPORT FOR CONFIG FILE
	if(config_path != ""){
		log.Println("reading config file")
	}

	// validate username
	if(username == ""){log.Fatalln("invalid username")}

	// set up max hashrate
	var max_hashrate int
	var err error
	if(max_hashrate_str != ""){
		switch max_hashrate_str{
			case "LOW_NANO":
				max_hashrate = LOW_NANO
				difficulty = "LOW"
				break
			
			case "LOW_ESP32":
				max_hashrate = LOW_ESP32
				difficulty = "LOW"
				break
			
			case "LOW_PI4":
				max_hashrate = LOW_PI4
				difficulty = "LOW"
				break
			default:
				max_hashrate, err = strconv.Atoi(max_hashrate_str)
				if(err != nil){
					log.Fatalln("error while setting max hashrate:", err)
				}
		}
	} else {
		max_hashrate = 0
	}
	log.Println("max_hashrate:", max_hashrate)
	log.Println("difficulty:", difficulty)

	miner := Worker{
		Username 		: username,
		Difficulty 		: difficulty,
		Software_name 	: software_name,
		Rig_name 		: rig_name,
		Xxhash			: xxhash,
		Job_type 		: "JOB",
		Max_hashrate	: max_hashrate,
		Max_rejected	: max_rejected,
	}

	miners := []Worker{}
	for i := 0; i <= workers; i++{
		miners = append(miners, miner)
		miners[i].Software_name = miners[i].Rig_name + " " + strconv.Itoa(i)
		log.Println("created worker " + strconv.Itoa(i))
		// setup xxhash
		if(xxhash){miners[i].Job_type = "JOBXX"; miners[i].Difficulty = "XXHASH"}
		// if(xxhash){log.Fatalln("xxHash is disabled")}
		if(i==workers){
			miners[i].Work()
		} else {
			time.Sleep(2* time.Second)
			go miners[i].Work()
		}
		log.Println("started worker " + strconv.Itoa(i))
	}
}