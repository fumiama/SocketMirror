//go:build windows
// +build windows

package main

import (
	"fmt"
	"sync"
	"syscall"
)

var (
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")

	setConsoleTextAttributeProc = kernel32DLL.NewProc("SetConsoleTextAttribute")
)

var pmu sync.Mutex

const (
	BLUE  = 0x0001 // FOREGROUND_BLUE
	GREEN = 0x0002 // FOREGROUND_GREEN
	RED   = 0x0004 // FOREGROUND_RED
	FITEN = 0x0008 // FOREGROUND_INTENSITY
)

func infoln(i int, title string, a ...any) {
	pmu.Lock()
	defer pmu.Unlock()
	setcolor(BLUE | FITEN)
	fmt.Print("INFO ")
	setcolor(i & 0xff)
	fmt.Print(title)
	setcolor(BLUE | GREEN | RED)
	fmt.Println(a...)
}
func erroln(i int, title string, a ...any) {
	pmu.Lock()
	defer pmu.Unlock()
	setcolor(RED | FITEN)
	fmt.Print("ERRO ")
	setcolor(i & 0xff)
	fmt.Print(title)
	setcolor(BLUE | GREEN | RED)
	fmt.Println(a...)
}
func infof(i int, title string, s string, a ...any) {
	pmu.Lock()
	defer pmu.Unlock()
	setcolor(BLUE | FITEN)
	fmt.Print("INFO ")
	setcolor(i & 0xff)
	fmt.Print(title)
	setcolor(BLUE | GREEN | RED)
	fmt.Printf(s, a...)
}

func setcolor(color int) {
	setConsoleTextAttribute(uintptr(syscall.Stdout), uint16(color))
}

// checkError evaluates the results of a Windows API call and returns the error if it failed.
func checkError(r1, r2 uintptr, err error) error {
	// Windows APIs return non-zero to indicate success
	if r1 != 0 {
		return nil
	}

	// Return the error if provided, otherwise default to EINVAL
	if err != nil {
		return err
	}
	return syscall.EINVAL
}

// SetConsoleTextAttribute sets the attributes of characters written to the
// console screen buffer by the WriteFile or WriteConsole function.
// See http://msdn.microsoft.com/en-us/library/windows/desktop/ms686047(v=vs.85).aspx.
func setConsoleTextAttribute(handle uintptr, attribute uint16) error {
	r1, r2, err := setConsoleTextAttributeProc.Call(handle, uintptr(attribute), 0)
	// use(attribute)
	return checkError(r1, r2, err)
}
