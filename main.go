package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	tm "github.com/buger/goterm"

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

// setupMainSignalHandler is for main to do cleanup
func setupMainSignalHandler() <-chan os.Signal {
	mc := make(chan os.Signal, 2)
	signal.Notify(mc, os.Interrupt, syscall.SIGTERM)
	return mc
}

func infoln(i int, title string, a ...any) {
	tm.Print(tm.Color("INFO ", tm.BLUE), tm.Color(title, i%8), " ", fmt.Sprintln(a...))
	tm.Flush()
}
func erroln(i int, title string, a ...any) {
	tm.Print(tm.Color("ERRO ", tm.RED), tm.Color(title, i%8), " ", fmt.Sprintln(a...))
	tm.Flush()
}
func infof(i int, title string, s string, a ...any) {
	tm.Print(tm.Color("INFO ", tm.BLUE), tm.Color(title, i%8), " ", fmt.Sprintf(s, a...))
	tm.Flush()
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
	mc := setupMainSignalHandler()
	if *c > 0 {
		cli := func(i int) {
			bm := client.NewBeamer(*addr)
			err := bm.Connect(*n, time.Second*time.Duration(*t))
			if err != nil {
				erroln(i, fmt.Sprintf("[%04d]", i), err)
				os.Exit(line())
			}
			hash, ch, msg := bm.Beam(int(*l), int(*r))
			go func() {
				for err := range ch {
					erroln(i, fmt.Sprintf("[%04d]", i), err)
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
				infof(i, fmt.Sprintf("[%04d]", i), "beam max: %dns, min: %dns, avg: %dns\n", max, min, sum/int64(*l))
			}()
			t0 := time.Now().UnixMilli()
			cnt, max, min, sum, err := bm.See(int(*l), int(*r), hash)
			if err != nil {
				erroln(i, fmt.Sprintf("[%04d]", i), err)
			}
			infof(i, fmt.Sprintf("[%04d]", i), "see max: %dns, min: %dns, avg: %dns\n", max, min, sum/int64(*l))
			t1 := time.Now().UnixMilli()
			totl := *l * *r
			delta := t1 - t0
			infof(i, fmt.Sprintf("[%04d]", i), "%d/%d succ/totl, speed: %0.2f B/s\n", cnt, *l, float64(totl)/float64(delta)/1000)
			bm.Close()
		}
		if *s {
			for i := 0; i < int(*c); i++ {
				go cli(i + 1)
			}
		} else {
			for i := 0; i < int(*c); i++ {
				go cli(i + 1)
			}
			<-mc
			return
		}
	}
	if *s {
		mr := server.NewMirror(*addr)
		if strings.Contains(*n, "tcp") {
			err := mr.Listen(*n)
			if err != nil {
				erroln(tm.WHITE, "[srv]", err)
				os.Exit(line())
			}
		}
		infoln(tm.WHITE, "[srv]", "start")
		ch, msg := mr.Reflect(int(*l))
		var (
			err      error
			str      string
			ok1, ok2 bool
		)
	lop:
		for {
			select {
			case err, ok1 = <-ch:
				erroln(tm.WHITE, "[srv]", err)
			case str, ok2 = <-msg:
				infoln(tm.WHITE, "[srv]", str)
			case <-mc:
				break lop
			}
			if !ok1 && !ok2 {
				break lop
			}
		}
		mr.Close()
		infoln(tm.WHITE, "[srv]", "terminate")
	}
}
