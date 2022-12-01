package main

import (
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
	mrerrch, mrmsgch := mr.Reflect(1024)
	bm := client.NewBeamer("127.0.0.1:45242")
	err = bm.Connect("tcp", time.Second)
	if err != nil {
		t.Fatal(err)
	}
	hash, bmerrch, bmmsgch := bm.Beam(1024, 1024)
	go func() {
		for err := range bmerrch {
			t.Log(err)
		}
	}()
	go func() {
		for err := range mrerrch {
			t.Log(err)
		}
	}()
	go func() {
		for err := range mrmsgch {
			t.Log(err)
		}
	}()
	go func() {
		for err := range bmmsgch {
			t.Log(err)
		}
	}()
	sc, _, _, _, err := bm.See(1024, 1024, hash)
	if err != nil {
		t.Fatal(err)
	}
	bm.Close()
	mr.Close()
	assert.Equal(t, 1024, sc)
}

func TestUdp(t *testing.T) {
	mr := server.NewMirror("127.0.0.1:45242")
	mrerrch, mrmsgch := mr.Reflect(1024)
	bm := client.NewBeamer("127.0.0.1:45242")
	err := bm.Connect("udp", time.Second)
	if err != nil {
		t.Fatal(err)
	}
	hash, bmerrch, bmmsgch := bm.Beam(1024, 1024)
	go func() {
		for err := range bmerrch {
			t.Log(err)
		}
	}()
	go func() {
		for err := range mrerrch {
			t.Log(err)
		}
	}()
	go func() {
		for err := range mrmsgch {
			t.Log(err)
		}
	}()
	go func() {
		for err := range bmmsgch {
			t.Log(err)
		}
	}()
	sc, _, _, _, err := bm.See(1024, 1024, hash)
	if err != nil {
		t.Fatal(err)
	}
	bm.Close()
	mr.Close()
	assert.Equal(t, 1024, sc)
}
