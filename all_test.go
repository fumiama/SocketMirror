package main

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/fumiama/SocketMirror/client"
	"github.com/fumiama/SocketMirror/server"
)

func TestTcp(t *testing.T) {
	mr := server.NewMirror("127.0.0.1:45242")
	err := mr.Listen("tcp")
	if err != nil {
		t.Fatal(err)
	}
	mrerrch := mr.Reflect(make([]byte, 1024))
	bm := client.NewBeamer("127.0.0.1:45242")
	err = bm.Connect("tcp", time.Second)
	if err != nil {
		t.Fatal(err)
	}
	hash, bmerrch, _ := bm.Beam(1024, 1024)
	sc, _, _, _, err := bm.See(1024, 1024, hash)
	if err != nil {
		t.Fatal(err)
	}
	defer bm.Close()
	for err := range bmerrch {
		t.Log(err)
	}
	err = mr.Close()
	if err != nil {
		t.Fatal(err)
	}
	for err := range mrerrch {
		t.Log(err)
	}
	assert.Equal(t, 1024, sc)
}

func TestUdp(t *testing.T) {
	mr := server.NewMirror("127.0.0.1:45242")
	mrerrch := mr.Reflect(make([]byte, 1024))
	bm := client.NewBeamer("127.0.0.1:45242")
	err := bm.Connect("udp", time.Second)
	if err != nil {
		t.Fatal(err)
	}
	hash, bmerrch, _ := bm.Beam(1024, 1024)
	sc, _, _, _, err := bm.See(1024, 1024, hash)
	if err != nil {
		t.Fatal(err)
	}
	bm.Close()
	err = <-bmerrch
	if err != nil {
		t.Fatal(err)
	}
	err = mr.Close()
	if err != nil {
		t.Fatal(err)
	}
	err = <-mrerrch
	if !strings.Contains(err.Error(), "use of closed network connection") {
		t.Fatal(err)
	}
	assert.Equal(t, 1024, sc)
}
