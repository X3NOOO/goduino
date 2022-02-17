package main

const (
	VERSION string = "0.1"
	GETPOOL_ADDR string = "https://server.duinocoin.com/getPool"
	// SERV_ADDR string = "51.15.127.80:2811"
	MAX_REJ int = 15
)

var AVAILABLE_PORTS = [...]string{"2811", "2812", "2813", "2814", "2815"}

var accepted int = 0
var rejected int = 0
