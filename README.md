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
```bash
$ go run main.go -c 1 -s 
checking...
client 1 send 1048576 bytes in 8 ms, speed: 131.072 bytes/s
client 1 total: 1024 batches success: 361 batches
$ go run main.go -c 1 -s 
checking...
client 1 send 1048576 bytes in 7 ms, speed: 149.7965714285714 bytes/s
client 1 total: 1024 batches success: 297 batches
$ go run main.go -c 1 -s 
checking...
client 1 send 1048576 bytes in 7 ms, speed: 149.7965714285714 bytes/s
client 1 total: 1024 batches success: 299 batches
$ go run main.go -c 1 -s 
checking...
client 1 send 1048576 bytes in 6 ms, speed: 174.76266666666666 bytes/s
client 1 total: 1024 batches success: 1024 batches
$ go run main.go -c 1 -s 
checking...
client 1 send 1048576 bytes in 7 ms, speed: 149.7965714285714 bytes/s
client 1 total: 1024 batches success: 1024 batches
$ go run main.go -c 1 -s 
checking...
client 1 send 1048576 bytes in 7 ms, speed: 149.7965714285714 bytes/s
client 1 total: 1024 batches success: 259 batches

$ go run main.go -c 10 -s 
checking...
checking...
checking...
checking...
checking...
checking...
checking...
checking...
checking...
checking...
client 2 send 1048576 bytes in 54 ms, speed: 19.418074074074074 bytes/s
client 2 total: 1024 batches success: 52 batches
client 8 send 1048576 bytes in 54 ms, speed: 19.418074074074074 bytes/s
client 8 total: 1024 batches success: 87 batches
client 6 send 1048576 bytes in 54 ms, speed: 19.418074074074074 bytes/s
client 6 total: 1024 batches success: 56 batches
client 10 send 1048576 bytes in 54 ms, speed: 19.418074074074074 bytes/s
client 10 total: 1024 batches success: 55 batches
client 9 send 1048576 bytes in 55 ms, speed: 19.065018181818182 bytes/s
client 9 total: 1024 batches success: 30 batches
client 3 send 1048576 bytes in 55 ms, speed: 19.065018181818182 bytes/s
client 3 total: 1024 batches success: 36 batches
client 1 send 1048576 bytes in 55 ms, speed: 19.065018181818182 bytes/s
client 1 total: 1024 batches success: 142 batches
client 4 send 1048576 bytes in 58 ms, speed: 18.07889655172414 bytes/s
client 4 total: 1024 batches success: 56 batches
client 5 send 1048576 bytes in 59 ms, speed: 17.772474576271186 bytes/s
client 5 total: 1024 batches success: 100 batches
client 7 send 1048576 bytes in 58 ms, speed: 18.07889655172414 bytes/s
client 7 total: 1024 batches success: 87 batches
```
