package server

import (
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/pkg/errors"
)

// Mirror ...
type Mirror struct {
	sync.Mutex
	addr string
	lsnr net.Listener
	conn []io.Closer
}

// NewMirror host:port
func NewMirror(address string) *Mirror {
	return &Mirror{
		addr: address,
	}
}

// Listen on network
func (mr *Mirror) Listen(network string) (err error) {
	mr.lsnr, err = net.Listen(network, mr.addr)
	return
}

// Reflect all data from peer
func (mr *Mirror) Reflect(bufsz int) (<-chan error, <-chan string) {
	ch := make(chan error, 2)
	msg := make(chan string, 256)
	if mr.lsnr != nil { // tcp
		go func() {
			wg := sync.WaitGroup{}
			for {
				conn, err := mr.lsnr.Accept()
				if err != nil {
					ch <- err
					break
				}
				msg <- fmt.Sprintf("accept: %v", conn.RemoteAddr())
				wg.Add(1)
				mr.Lock()
				mr.conn = append(mr.conn, conn)
				mr.Unlock()
				go func() {
					buf := make([]byte, bufsz)
					defer func() {
						mr.Lock()
						defer mr.Unlock()
						conn.Close()
						for i, c := range mr.conn {
							if c == conn {
								mr.conn = append(mr.conn[:i], mr.conn[i+1:]...)
								return
							}
						}
					}()
					defer wg.Done()
					for {
						n, err := conn.Read(buf)
						if err != nil {
							if err != io.EOF {
								ch <- errors.Wrapf(err, "read from %v", conn.RemoteAddr())
							}
							break
						}
						_, err = conn.Write(buf[:n])
						if err != nil {
							if err != io.EOF {
								ch <- errors.Wrapf(err, "write to %v", conn.RemoteAddr())
							}
							break
						}
					}
					msg <- fmt.Sprintf("close: %v", conn.RemoteAddr())
				}()
			}
			wg.Wait()
			close(ch)
		}()
		return ch, msg
	}
	// udp
	go func() {
		defer close(ch)
		conn, err := net.ListenPacket("udp", mr.addr)
		if err != nil {
			ch <- err
			return
		}
		mr.Lock()
		mr.conn = append(mr.conn, conn)
		mr.Unlock()
		buf := make([]byte, bufsz)
		for {
			n, addr, err := conn.ReadFrom(buf)
			if err != nil {
				ch <- errors.Wrapf(err, "read from %v", addr)
				break
			}
			_, err = conn.WriteTo(buf[:n], addr)
			if err != nil {
				ch <- errors.Wrapf(err, "write to %v", addr)
			}
		}
	}()
	return ch, msg
}

func (mr *Mirror) Close() {
	mr.Lock()
	defer mr.Unlock()
	for _, conn := range mr.conn {
		_ = conn.Close()
	}
	mr.conn = mr.conn[:0]
	if mr.lsnr != nil {
		mr.lsnr.Close()
		mr.lsnr = nil
	}
}
