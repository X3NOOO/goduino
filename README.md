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
  -h
        display this help message
```

## donation

- XMR: `49F3GknYgs7cRfMJghrd9dHZKe63Z6Y3aJKPecDKqLRje5YebzWvz3VWsTa8e8Sk92G7WJEsyp8L1VEeNxmdj2vZNJSACo1`
- DOGE: `DFYc29EsSuSbyLndGrKBGoC2usHRUqiiXb`
- BTC: `bc1q08p6wd86806uf2cj95j4pcgl584jvaqkhs37pp`
- ETH: `0x84FfD8524a66505344A1cbfC3212392Db5b2474d`
- LTC: `Lew3VmzbkaxzoYG3jNHf263oEDMrQ3ecN1`

## todo

[X] add xxhash support
[ ] add hashrate limiter
[ ] user-friendly logs
