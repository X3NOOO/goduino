package main

import (
	// "fmt"
	// "encoding/json"
	// "io/ioutil"
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"github.com/cespare/xxhash"
	// "net/http"
	"strconv"
	// "strconv"
)

func getFastestServer() (string, error){
	// download data from server in GETPOOL_ADDR
	rawresp, err := http.Get(GETPOOL_ADDR)
	if(err != nil){
	return "", err
	}
	defer rawresp.Body.Close()
	// resp, err := net.Dial("tcp", GETPOOL_ADDR)
	log.Println("raw response:", rawresp.Body)

	// parse data
	resp, err := ioutil.ReadAll(rawresp.Body)
	if(err != nil){
	return "", err
	}
	respstr := string(resp)
	log.Println("response:", respstr)

	return respstr, err
}

type ServerDetails struct {
	Ip 		string
	Name 	string
	Port 	int
	Server 	string
	Success bool
}

type Worker struct {
	Username 		string
	Difficulty 		string
	Software_name 	string
	Rig_name 		string
	Xxhash			bool
	Job_type		string
}

func (self Worker) Work() {
	var username, difficulty, software_name, rig_name, job_type string = self.Username, self.Difficulty, self.Software_name, self.Rig_name, self.Job_type
	var xxhash_enable bool = self.Xxhash
	// Get server details from GETPOOL_ADDR
	log.Println("getting fastest server")
	pool_address, err := getFastestServer()
	if(err != nil){
		log.Fatalln(err)
	}
	
	var server_details ServerDetails
	json.Unmarshal([]byte(pool_address), &server_details)
	
	log.Println("server_details:", server_details)

	// Start tcp connection
	log.Println("starting tcp connection")
	ALL_PORTS := []string{strconv.Itoa(server_details.Port)}
	ALL_PORTS = append(ALL_PORTS, AVAILABLE_PORTS[:]...) 
	log.Println("ALL_PORTS:", ALL_PORTS)
	var conn net.Conn
	for i := 0; i <= len(ALL_PORTS); i++ {
		conn, err = net.Dial("tcp", server_details.Ip + ":" + ALL_PORTS[i])
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

	// start mining
	for {
		// get a job
		log.Println("getting job")
		log.Println("job request:", job_type + "," + username + "," + difficulty)
		_, err = conn.Write([]byte(job_type + "," + username + "," + difficulty))
		if(err != nil){
			log.Println("error while getting a job:", err)
			continue
		}
		job_buff := make([]byte, 1024)
		_, err = conn.Read(job_buff)
		if(err != nil){
			if(xxhash_enable){
				log.Fatalln("xxHash is currently disabled")
			}
			log.Println("error while reading a job:", err)
			continue
		}
		// parse job
		job_buff = bytes.Trim(job_buff, "\x00")
		job := strings.Split(strings.TrimSpace(string(job_buff)), ",")
		log.Println("job:", job)
		pref_job := job[0]
		target_job := job[1]
		diff_job, _ := strconv.Atoi(job[2])

		// brute force hash
		var hash string
		for i := 0; i <= diff_job*100; i++{
			if(xxhash_enable){
				h := xxhash.New()
				h.Write([]byte(pref_job + strconv.Itoa(i)))
				hash = hex.EncodeToString(h.Sum(nil))
			} else {
				h := sha1.New()
				h.Write([]byte(pref_job + strconv.Itoa(i)))
				hash = hex.EncodeToString(h.Sum(nil))
			}

			if(hash == target_job){
				log.Println("guessed hash " + hash + " on " + strconv.Itoa(i))
				log.Println("sending to server: " + strconv.Itoa(i) + "," + strconv.Itoa(diff_job) + "," + software_name + "," + rig_name)
				_, err = conn.Write([]byte(strconv.Itoa(i) + "," + strconv.Itoa(diff_job) + "," + software_name + "," + rig_name))
				if(err != nil){
					log.Println("error while sending hash result")
					break
				}
				resp_buff := make([]byte, 32)
				_, err = conn.Read(resp_buff)
				if(err != nil){
					log.Println("error while revicing result")
					break
				}
				resp_buff = bytes.Trim(resp_buff, "\x00")
				resp := strings.TrimSpace((string(resp_buff)))
				log.Println("server's respond: " + resp)

				if(resp == "GOOD" || resp == "BLOCK"){
					log.Println("guessed correct hash on " + strconv.Itoa(i))
					accepted++
				} else if(resp == "BAD"){
					log.Println("guessed invalid hash on " + strconv.Itoa(i))
					rejected++
				} else if(resp == "INVU"){
					log.Fatalln("invalid username")
				}
			}
		}
		if(rejected >= MAX_REJ){
			log.Fatalln("reached max rejected hashes, exiting")
		}
	}
}