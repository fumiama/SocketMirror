package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"sync"
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
	c := flag.Uint("c", 0, "client count")
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
	if !*s && *c == 0 {
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
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	if *s {
		mr := server.NewMirror(*addr)
		if strings.Contains(*n, "tcp") {
			err := mr.Listen(*n)
			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(line())
			}
		}
		wg.Add(1)
		ch := mr.Reflect(int(*l))
		go func() {
			for err := range ch {
				fmt.Println("ERROR:", err)
			}
			wg.Done()
		}()
	}
	if *c > 0 {
		wg.Add(int(*c))
		for i := 0; i < int(*c); i++ {
			go func(i int) {
				bm := client.NewBeamer(*addr)
				err := bm.Connect(*n, time.Second*time.Duration(*t))
				if err != nil {
					fmt.Println("ERROR:", err)
					os.Exit(line())
				}
				hash, ch, msg := bm.Beam(int(*l), int(*r))
				go func() {
					for err := range ch {
						fmt.Println("ERROR:", err)
					}
				}()
				go func() {
					max := int64(0)
					min := int64(math.MaxInt64)
					sum := int64(0)
					for d := range msg {
						if d > max {
							max = d
						}
						if d < min {
							min = d
						}
						sum += d
					}
					fmt.Printf("%04d -> beam max: %dns, min: %dns, avg: %dns\n", i, max, min, sum/int64(*l))
				}()
				t0 := time.Now().UnixMilli()
				cnt, max, min, sum, err := bm.See(int(*l), int(*r), hash)
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				fmt.Printf("%04d -> see max: %dns, min: %dns, avg: %dns\n", i, max, min, sum/int64(*l))
				t1 := time.Now().UnixMilli()
				totl := *l * *r
				delta := t1 - t0
				fmt.Printf("%04d -> %d/%d succ/totl, speed: %0.2f B/s\n", i, cnt, *l, float64(totl)/float64(delta)/1000)
				wg.Done()
			}(i + 1)
		}
	}
	wg.Wait()
}
