# goduino

goduino is miner for DUCO (<https://duinocoin.com/>) that support emulating your other devices to optimize hashrate

## installation

from binary:

1. download binary for your platform from release

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

- XMR: `49F3GknYgs7cRfMJghrd9dHZKe63Z6Y3aJKPecDKqLRje5YebzWvz3VWsTa8e8Sk92G7WJEsyp8L1VEeNxmdj2vZNJSACo1`
- DOGE: `DFYc29EsSuSbyLndGrKBGoC2usHRUqiiXb`
- BTC: `bc1q08p6wd86806uf2cj95j4pcgl584jvaqkhs37pp`
- ETH: `0x84FfD8524a66505344A1cbfC3212392Db5b2474d`
- LTC: `Lew3VmzbkaxzoYG3jNHf263oEDMrQ3ecN1`

## todo

- [X] add xxhash support
- [X] add hashrate limiter
- [X] user-friendly logs
- [X] update readme.md
- [ ] support for config file
