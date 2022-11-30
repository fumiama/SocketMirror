package client

import (
	"hash/crc64"
	"io"
	"math/rand"
	"net"
	"time"
)

// Beamer ...
type Beamer struct {
	addr string
	conn net.Conn
}

// NewBeamer host:port
func NewBeamer(address string) *Beamer {
	return &Beamer{
		addr: address,
	}
}

// Connect to address in timeout
func (bm *Beamer) Connect(network string, timeout time.Duration) (err error) {
	bm.conn, err = net.DialTimeout(network, bm.addr, timeout)
	return
}

// Beam send batchsize bytes random data by batchcount times
//
//go:nosplit
func (bm *Beamer) Beam(batchsize, batchcount int) (uint64, <-chan error) {
	data := make([]byte, batchsize)
	_, _ = rand.Read(data)
	h := crc64.New(crc64.MakeTable(crc64.ISO))
	h.Write(data)
	hash := h.Sum64()
	ch := make(chan error, 2)
	go func() {
		defer close(ch)
		for batchcount > 0 {
			_, err := bm.conn.Write(data)
			if err != nil {
				ch <- err
				return
			}
			batchcount--
		}
	}()
	return hash, ch
}

// See whether all data have been reflected by peer
//
//go:nosplit
func (bm *Beamer) See(batchsize, batchcount int, hash uint64) (successcount int, err error) {
	data := make([]byte, batchsize)
	h := crc64.New(crc64.MakeTable(crc64.ISO))
	for batchcount > 0 {
		_, err = io.ReadFull(bm.conn, data)
		if err != nil {
			return
		}
		h.Write(data)
		datahash := h.Sum64()
		if datahash == hash {
			successcount++
		}
		h.Reset()
		batchcount--
	}
	return
}

// Close ...
func (bm *Beamer) Close() (err error) {
	return bm.conn.Close()
}
