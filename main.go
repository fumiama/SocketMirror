package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/fumiama/SocketMirror/client"
	"github.com/fumiama/SocketMirror/server"
)

func line() int {
	_, _, fileLine, ok := runtime.Caller(1)
	if ok {
		return fileLine
	}
	return -1
}

func main() {
	n := flag.String("n", "tcp", "network")
	addr := flag.String("a", "127.0.0.1:8888", "server addr")
	s := flag.Bool("s", false, "server mode")
	c := flag.Bool("c", false, "client mode")
	l := flag.Uint("l", 1024, "batch size (length per write)")
	r := flag.Uint("r", 1024, "batch count (repeat times)")
	t := flag.Uint("t", 4, "connect timeout (s)")
	h := flag.Bool("h", false, "display this help")
	flag.Parse()
	if *h {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	}
	if !*s && !*c {
		fmt.Println("ERROR: -s and -c can not be disabled both")
		os.Exit(line())
	}
	if !strings.Contains(*n, "udp") && !strings.Contains(*n, "tcp") {
		fmt.Println("ERROR: network must be tcp or udp")
		os.Exit(line())
	}
	if *l == 0 {
		fmt.Println("ERROR: zero batch size")
		os.Exit(line())
	}
	if *r == 0 {
		fmt.Println("ERROR: zero batch count")
		os.Exit(line())
	}
	if *s {
		mr := server.NewMirror(*addr)
		if strings.Contains(*n, "tcp") {
			err := mr.Listen(*n)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(line())
			}
		}
		ch := mr.Reflect(make([]byte, *l))
		go func() {
			for err := range ch {
				fmt.Println("ERROR:", err)
			}
		}()
	}
	if *c {
		bm := client.NewBeamer(*addr)
		err := bm.Connect(*n, time.Second*time.Duration(*t))
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(line())
		}
		hash, ch := bm.Beam(int(*l), int(*r))
		go func() {
			for err := range ch {
				fmt.Println("ERROR:", err)
			}
		}()
		fmt.Println("checking...")
		t0 := time.Now().UnixMilli()
		cnt, err := bm.See(int(*l), int(*r), hash)
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		t1 := time.Now().UnixMilli()
		totl := *l * *r
		delta := t1 - t0
		fmt.Println("send", totl, "bytes in", delta, "ms, speed:", float64(totl)/float64(delta)/1000, "bytes/s")
		fmt.Println("total:", *l, "batches", "success:", cnt, "batches")
	}
}