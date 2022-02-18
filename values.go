package main

const (
	VERSION string = "0.1"
	GETPOOL_ADDR string = "https://server.duinocoin.com/getPool"
	// SERV_ADDR string = "51.15.127.80:2811"

	// prefixes
	K int = 1000
	M int = K*K

	// boards hashrate
	LOW_NANO 	int = 200
	LOW_ESP32 	int = 34*K
	LOW_PI4 	int = 742*K
)

var AVAILABLE_PORTS = [...]string{"2811", "2812", "2813", "2814", "2815"}

var accepted int = 0
var rejected int = 0
