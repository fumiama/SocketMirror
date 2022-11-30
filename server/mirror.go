package server

import (
	"io"
	"net"
	"sync"
)

// Mirror ...
type Mirror struct {
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
func (mr *Mirror) Reflect(buf []byte) <-chan error {
	ch := make(chan error, 2)
	if mr.lsnr != nil { // tcp
		go func() {
			wg := sync.WaitGroup{}
			for {
				conn, err := mr.lsnr.Accept()
				if err != nil {
					ch <- err
					break
				}
				wg.Add(1)
				mr.conn = append(mr.conn, conn)
				go func() {
					defer wg.Done()
					for {
						n, err := conn.Read(buf)
						if err != nil {
							ch <- err
							return
						}
						_, err = conn.Write(buf[:n])
						if err != nil {
							ch <- err
						}
					}
				}()
			}
			wg.Wait()
			close(ch)
		}()
		return ch
	}
	// udp
	go func() {
		defer close(ch)
		conn, err := net.ListenPacket("udp", mr.addr)
		mr.conn = append(mr.conn, conn)
		if err != nil {
			ch <- err
			return
		}
		for {
			n, addr, err := conn.ReadFrom(buf)
			if err != nil {
				ch <- err
				return
			}
			_, err = conn.WriteTo(buf[:n], addr)
			if err != nil {
				ch <- err
			}
		}
	}()
	return ch
}

func (mr *Mirror) Close() error {
	for _, conn := range mr.conn {
		_ = conn.Close()
	}
	if mr.lsnr != nil {
		return mr.lsnr.Close()
	}
	return nil
}
