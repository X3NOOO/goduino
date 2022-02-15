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
	"syscall"
	// "time"
)

// usage: 	goduino mine [optional: -workers=how_many_devices] [optional: -rig_name=rig_name] [optional: -software_name=software_name] [optional: -debug true] -name=username -diff=difficulty (LOW (for web miners and low-spec devices), MEDIUM (pc), NET (network diff - pc), EXTREME (officially not used anywhere, but you can use it to mine or gpus - howewer this program currently don't support mining on gpus))
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


	flag.Parse()

	var workers int = *flag_workers
	var rig_name string = *flag_rig_name
	var software_name string = *flag_software_name
	var debug bool = *flag_debug
	var username string = *flag_name
	var difficulty string = *flag_diff
	var config_path string = *flag_config
	
	// Set up logging
	if(!debug){log.SetOutput(ioutil.Discard)}
	
	log.Println("workers:", workers)
	log.Println("rig_name:", rig_name)
	log.Println("software_name:", software_name)
	log.Println("debug:", debug)
	log.Println("username:", username)
	log.Println("difficulty:", difficulty)
	log.Println("config_path:", config_path)
	log.Println("VERSION:", VERSION)
	log.Println("AVAILABLE_PORTS:", AVAILABLE_PORTS)

	//read config
	if(config_path != ""){
		log.Println("reading config file")
	}

	// validate username
	if(username == ""){log.Fatalln("invalid username")}

	work()
}