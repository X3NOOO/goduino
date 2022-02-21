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
	l := Logger {}
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
	flag_max_rejected := flag.Int("max_rejected", 0, "after how many rejected hashes you want to exit (default: disabled)")


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
	
	l.error("That's how error looks like")
	l.warning("That's how warning looks like")
	l.info("That's how information looks like")
	l.guessed("That's how information about guessed hash looks like")
	l.hashrate("That's how hashrate message looks like")
	l.job([]string{"That's how job message looks like"})
	l.server("That's how server message looks like")

	log.Println("workers:", workers)
	log.Println("rig_name:", rig_name)
	log.Println("software_name:", software_name)
	log.Println("debug:", debug)
	log.Println("username:", username)
	log.Println("config_path:", config_path) 
	log.Println("max_hashrate_str:", max_hashrate_str)
	log.Println("VERSION:", VERSION)

	l.info("workers: " + strconv.Itoa(workers))
	l.info("Rig name: " + rig_name)
	l.info("Software name: " + software_name)
	l.info("Enabled debug: " + strconv.FormatBool(debug))
	l.info("Username: " + username)
	l.info("Config path: " + config_path) 
	l.info("Version: " + VERSION)

	//read config //TODO ADD SUPPORT FOR CONFIG FILE
	if(config_path != ""){
		log.Println("reading config file")
		l.error("Config file isn't supported yet")
		log.Fatalln("config file isnt supported yet")
	}

	// validate username
	if(username == ""){l.error("Invalid username");log.Fatalln("invalid username")}


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
					l.error("While setting max hashrate: ", err)
				}
		}
	} else {
		max_hashrate = 0
	}
	l.info("Max hashrate: " + strconv.Itoa(max_hashrate))
	l.info("Difficulty: " + difficulty)
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
		l.info("Created worker " + strconv.Itoa(i))
		// setup xxhash
		if(xxhash){miners[i].Job_type = "JOBXX"; miners[i].Difficulty = "XXHASH"}
		// if(xxhash){log.Fatalln("xxHash is disabled")}
		if(i==workers){
			log.Println("Starting worker " + strconv.Itoa(i))
			l.info("Starting worker " + strconv.Itoa(i))
			miners[i].Work()
		} else {
			time.Sleep(2 * time.Second)
			log.Println("Starting worker " + strconv.Itoa(i))
			l.info("Starting worker " + strconv.Itoa(i))
			go miners[i].Work()
		}
	}
}