//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"sync"

	tm "github.com/buger/goterm"
)

var pmu sync.Mutex

func infoln(i int, title string, a ...any) {
	pmu.Lock()
	defer pmu.Unlock()
	tm.Print(tm.Color("INFO ", tm.BLUE), tm.Color(title, i%8), " ", fmt.Sprintln(a...))
	tm.Flush()
}
func erroln(i int, title string, a ...any) {
	pmu.Lock()
	defer pmu.Unlock()
	tm.Print(tm.Color("ERRO ", tm.RED), tm.Color(title, i%8), " ", fmt.Sprintln(a...))
	tm.Flush()
}
func infof(i int, title string, s string, a ...any) {
	pmu.Lock()
	defer pmu.Unlock()
	tm.Print(tm.Color("INFO ", tm.BLUE), tm.Color(title, i%8), " ", fmt.Sprintf(s, a...))
	tm.Flush()
}
