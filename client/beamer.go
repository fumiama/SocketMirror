package client

import (
	"hash/crc64"
	"io"
	"math"
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
func (bm *Beamer) Beam(batchsize, batchcount int) (uint64, <-chan error, <-chan int64) {
	data := make([]byte, batchsize)
	_, _ = rand.Read(data)
	h := crc64.New(crc64.MakeTable(crc64.ISO))
	h.Write(data)
	hash := h.Sum64()
	ch := make(chan error, 2)
	msg := make(chan int64, batchcount)
	go func() {
		defer close(ch)
		defer close(msg)
		for batchcount > 0 {
			var err error
			cnt := 0
			n := 0
			t0 := time.Now().UnixNano()
			for cnt < batchsize {
				n, err = bm.conn.Write(data[cnt:])
				if err != nil {
					ch <- err
					return
				}
				cnt += n
			}
			t1 := time.Now().UnixNano()
			msg <- t1 - t0
			batchcount--
		}
	}()
	return hash, ch, msg
}

// See whether all data have been reflected by peer
//
//go:nosplit
func (bm *Beamer) See(batchsize, batchcount int, hash uint64) (successcount int, max, min, sum int64, err error) {
	data := make([]byte, batchsize)
	h := crc64.New(crc64.MakeTable(crc64.ISO))
	min = int64(math.MaxInt64)
	for batchcount > 0 {
		t0 := time.Now().UnixNano()
		_, err = io.ReadFull(bm.conn, data)
		t1 := time.Now().UnixNano()
		d := t1 - t0
		if d > max {
			max = d
		}
		if d < min {
			min = d
		}
		sum += d
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
