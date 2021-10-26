# yisheng-txs

scan eth-like blockchain and filter specified transactions

## building

```shell
make
```

this will generate a binary file `./build/bin/yisheng-txsa`,  
and an example config file of `scanswap` subcommand [config-example.toml](https://github.com/jowenshaw/yisheng-txs/blob/master/params/config-example.toml)

## help

#### yisheng-txs

```shell
./build/bin/yisheng-txs -h
```

```text
NAME:
   yisheng-txs - scan eth like blockchain

USAGE:
   yisheng-txs [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   scanswap  scan cross chain swaps
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbosity value  0:panic, 1:fatal, 2:error, 3:warn, 4:info, 5:debug, 6:trace (default: 4)
   --json             output log in json format (default: false)
   --color            output log in color text format (default: true)
   --help, -h         show help (default: false)
   --version, -v      print the version (default: false)
```

#### yisheng-txs scanswap

```shell
./build/bin/yisheng-txs scanswap -h
```

```text
NAME:
   yisheng-txs scanswap - scan cross chain swaps

USAGE:
   yisheng-txs scanswap [command options]

DESCRIPTION:
   scan cross chain swaps

OPTIONS:
   --config value, -c value  Specify config file
   --gateway value           gateway URL to connect
   --scanReceipt             scan transaction receipt instead of transaction (default: false)
   --start value             start height (start inclusive) (default: 0)
   --end value               end height (end exclusive) (default: 0)
   --stable value            stable height (default: 5)
   --jobs value              number of jobs (default: 4)
   --help, -h                show help (default: false)
```
