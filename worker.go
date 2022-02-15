package main

import (
	// "fmt"
	// "encoding/json"
	// "io/ioutil"
	"log"
	"net"
	// "net/http"
	"strconv"
	// "strconv"
)

// func getFastestServer() (string, error){
	// rawresp, err := http.Get(GETPOOL_ADDR)
	// if(err != nil){
		// return "", err
	// }
	// defer rawresp.Body.Close()
	// // resp, err := net.Dial("tcp", GETPOOL_ADDR)
	// log.Println("raw response:", rawresp.Body)
// 
	// resp, err := ioutil.ReadAll(rawresp.Body)
	// if(err != nil){
		// return "", err
	// }
	// respstr := string(resp)
	// log.Println("response:", respstr)
// 
	// return respstr, err
// }
// 
// type ServerDetails struct {
	// Ip 		string
	// Name 	string
	// Port 	int
	// Server 	string
	// Success bool
// }

func work(){
	// Get server details from GETPOOL_ADDR
	// log.Println("getting fastest server")
	// pool_address, err := getFastestServer()
	// if(err != nil){
		// log.Fatalln(err)
	// }
	// 
	// var server_details ServerDetails
	// json.Unmarshal([]byte(pool_address), &server_details)
	// 
	// log.Println("server_details:", server_details)
	// Start tcp connection
	log.Println("starting tcp connection")
	var conn net.Conn
	var err error
	for i := 0; i <= len(AVAILABLE_PORTS); i++{
		conn, err = net.Dial("tcp", POOL_ADDR + ":" + AVAILABLE_PORTS[i])
		//TODO uncomment line above for release and comment line below
		// conn, err = net.Dial("tcp", "51.15.127.80:2813")
		if(err == nil){
			log.Println("connected on port " + AVAILABLE_PORTS[i])
			break
		} else if(i < len(AVAILABLE_PORTS)){
			log.Println("error while connecting on port " + AVAILABLE_PORTS[i] + ":", err)
		} else {
			log.Fatalln("error on connecting on all available ports")
		}
	}
	log.Println("established tcp connection")

	server_version := make([]byte, 8)
	n, err := conn.Read(server_version)
	if(err != nil){
		log.Fatalln(err)
	}
	log.Println("Received " + strconv.Itoa(n) + " bytes")
	log.Println("Server version:", string(server_version))

	

}