# SocketMirror
Send random bytes to remote port and/or reflect all data to sender

## Usage
```bash
  -a string
        server addr (default "127.0.0.1:8888")
  -c    client mode
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
```bash
$ go run main.go -c -s 
checking...
send 1048576 bytes in -1669813308538372639 ms, speed: -6.27960020822833e-16 bytes/s
total: 1024 batches success: 1024 batches
$ go run main.go -c -s 
checking...
send 1048576 bytes in 6 ms, speed: 174.76266666666666 bytes/s
total: 1024 batches success: 1024 batches
$ go run main.go -c -s 
checking...
send 1048576 bytes in 7 ms, speed: 149.7965714285714 bytes/s
total: 1024 batches success: 387 batches
$ go run main.go -c -s 
checking...
send 1048576 bytes in 8 ms, speed: 131.072 bytes/s
total: 1024 batches success: 250 batches
$ go run main.go -c -s 
checking...
send 1048576 bytes in 7 ms, speed: 149.7965714285714 bytes/s
total: 1024 batches success: 269 batches
$ go run main.go -c -s 
checking...
send 1048576 bytes in 8 ms, speed: 131.072 bytes/s
total: 1024 batches success: 295 batches
$ go run main.go -c -s 
checking...
send 1048576 bytes in 7 ms, speed: 149.7965714285714 bytes/s
total: 1024 batches success: 281 batches
```
