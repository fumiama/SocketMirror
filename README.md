# SocketMirror
Send random bytes to remote port and/or reflect all data to sender

## Usage
```bash
  -a string
        server addr (default "127.0.0.1:8888")
  -c uint
        client count
  -h    display this help
  -l uint
        batch size (length per write) (default 1024)
  -n string
        network (default "tcp")
  -r uint
        batch count (repeat times) (default 1024)
  -s    server mode
  -t uint
        connect timeout (s) (default 4)
```

## Example
### 1 client
```bash
$ go run main.go -c 1 -s 
0001 -> 1024/1024 succ/totl, speed: 149.80 B/s
```
### 10 clients
```bash
$ go run main.go -c 10 -s 
0009 -> 112/1024 succ/totl, speed: 22.80 B/s
0007 -> 96/1024 succ/totl, speed: 21.85 B/s
0005 -> 123/1024 succ/totl, speed: 21.85 B/s
0010 -> 94/1024 succ/totl, speed: 21.40 B/s
0008 -> 91/1024 succ/totl, speed: 21.40 B/s
0002 -> 98/1024 succ/totl, speed: 20.97 B/s
0001 -> 85/1024 succ/totl, speed: 20.97 B/s
0004 -> 96/1024 succ/totl, speed: 19.42 B/s
0003 -> 215/1024 succ/totl, speed: 19.42 B/s
0006 -> 63/1024 succ/totl, speed: 19.42 B/s
```
### 64 clitens
```bash
$ go run main.go -c 64 -s 
0008 -> 102/1024 succ/totl, speed: 3.13 B/s
0001 -> 123/1024 succ/totl, speed: 3.11 B/s
0064 -> 128/1024 succ/totl, speed: 3.06 B/s
0052 -> 112/1024 succ/totl, speed: 2.83 B/s
0007 -> 63/1024 succ/totl, speed: 2.76 B/s
0010 -> 87/1024 succ/totl, speed: 2.67 B/s
0031 -> 30/1024 succ/totl, speed: 2.62 B/s
0028 -> 55/1024 succ/totl, speed: 2.60 B/s
0040 -> 153/1024 succ/totl, speed: 2.59 B/s
0041 -> 32/1024 succ/totl, speed: 2.58 B/s
0022 -> 21/1024 succ/totl, speed: 2.56 B/s
0032 -> 90/1024 succ/totl, speed: 2.54 B/s
0029 -> 36/1024 succ/totl, speed: 2.52 B/s
0055 -> 45/1024 succ/totl, speed: 2.49 B/s
0013 -> 37/1024 succ/totl, speed: 2.44 B/s
0043 -> 77/1024 succ/totl, speed: 2.40 B/s
0033 -> 196/1024 succ/totl, speed: 2.35 B/s
0018 -> 39/1024 succ/totl, speed: 2.27 B/s
0034 -> 66/1024 succ/totl, speed: 2.26 B/s
0057 -> 67/1024 succ/totl, speed: 2.26 B/s
0044 -> 59/1024 succ/totl, speed: 2.25 B/s
0025 -> 203/1024 succ/totl, speed: 2.25 B/s
0019 -> 254/1024 succ/totl, speed: 2.25 B/s
0058 -> 141/1024 succ/totl, speed: 2.24 B/s
0035 -> 248/1024 succ/totl, speed: 2.23 B/s
0045 -> 224/1024 succ/totl, speed: 2.23 B/s
0026 -> 81/1024 succ/totl, speed: 2.20 B/s
0020 -> 81/1024 succ/totl, speed: 2.14 B/s
0059 -> 61/1024 succ/totl, speed: 2.14 B/s
0011 -> 181/1024 succ/totl, speed: 2.11 B/s
0014 -> 24/1024 succ/totl, speed: 2.10 B/s
0030 -> 36/1024 succ/totl, speed: 2.10 B/s
0054 -> 44/1024 succ/totl, speed: 2.09 B/s
0009 -> 30/1024 succ/totl, speed: 2.09 B/s
0053 -> 115/1024 succ/totl, speed: 2.08 B/s
0016 -> 13/1024 succ/totl, speed: 2.08 B/s
0012 -> 32/1024 succ/totl, speed: 2.08 B/s
0047 -> 59/1024 succ/totl, speed: 2.08 B/s
0037 -> 53/1024 succ/totl, speed: 2.08 B/s
0050 -> 192/1024 succ/totl, speed: 2.10 B/s
0061 -> 55/1024 succ/totl, speed: 2.08 B/s
0048 -> 12/1024 succ/totl, speed: 2.08 B/s
0060 -> 70/1024 succ/totl, speed: 2.08 B/s
0038 -> 40/1024 succ/totl, speed: 2.08 B/s
0062 -> 37/1024 succ/totl, speed: 2.08 B/s
0039 -> 38/1024 succ/totl, speed: 2.08 B/s
0024 -> 59/1024 succ/totl, speed: 2.07 B/s
0063 -> 53/1024 succ/totl, speed: 2.08 B/s
0021 -> 67/1024 succ/totl, speed: 2.08 B/s
0049 -> 77/1024 succ/totl, speed: 2.08 B/s
0002 -> 297/1024 succ/totl, speed: 2.31 B/s
0004 -> 316/1024 succ/totl, speed: 2.31 B/s
0005 -> 252/1024 succ/totl, speed: 2.31 B/s
0006 -> 281/1024 succ/totl, speed: 2.31 B/s
0056 -> 76/1024 succ/totl, speed: 2.04 B/s
0017 -> 82/1024 succ/totl, speed: 2.04 B/s
0042 -> 82/1024 succ/totl, speed: 2.04 B/s
0023 -> 90/1024 succ/totl, speed: 2.04 B/s
0036 -> 85/1024 succ/totl, speed: 2.01 B/s
0046 -> 60/1024 succ/totl, speed: 2.01 B/s
0027 -> 61/1024 succ/totl, speed: 2.01 B/s
0015 -> 40/1024 succ/totl, speed: 2.00 B/s
0051 -> 747/1024 succ/totl, speed: 2.07 B/s
0003 -> 361/1024 succ/totl, speed: 1.90 B/s
```
