# goduino

goduino is miner for DUCO (<https://duinocoin.com/>) that support emulating other devices to optimize hashrate

# WARNING

remember of duino [TOS](https://duinocoin.com/legal)

## installation

from binary:

1. download binary for your platform from release page

from source:

1. `git clone https://github.com/X3NOOO/goduino`
2. `cd goduino`
3. `make release`

## usage

```raw
Usage of goduino:
  -config string
    	enter path to your configuration file
  -debug
    	enable debug logs
  -diff string
    	difficulty that you want to mine on (default "MEDIUM")
  -max_hashrate string
    	do you want to limit hashrate to some value (supported options: LOW_NANO, LOW_ESP32, LOW_PI4; or: value in H/s eg. 720 will mean 720H/s)
  -max_rejected int
    	after how many rejected hashes you want to exit (default: disabled)
  -name string
    	your username
  -rig_name string
    	custom rig name (default "goduino 0.1 worker")
  -software_name string
    	custom mining software name (default "goduino 0.1")
  -workers int
    	how many devices you want to use (default 1)
  -xxhash
    	do you want to use xxhash instead of sha1
```

## donation

- XMR: `
8BrqGJBJ9cAKWLhaZws37AbCZtKVg2cfq8JpNr6GmeuZYZUUHLgn2L4PLxg1eZHvzMLNncyYpduVWHb8X49qx8vmAL5oanL`
- DOGE: `DFYc29EsSuSbyLndGrKBGoC2usHRUqiiXb`
- BTC: `bc1q522kupdywc9zne68whttwcutwq2lfuzat825m4`
- ETH: `0xd3A2F08A920E581aeDD9edF0990fACf612A85A22`
- LTC: `ltc1qwva2zxvp77gehpazmmlus7377pypmraqtj8n8f`

## todo

- [X] add xxhash support
- [X] add hashrate limiter
- [X] user-friendly logs
- [X] update readme.md
- [ ] support for config file
