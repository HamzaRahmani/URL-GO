package tests

import (
	"net"
	"testing"
	"time"
)

// GetFreeTCPPort asks the kernel for a free open port that is ready to use.
func GetFreeTCPPort(t *testing.T) (port int, err error) {
	t.Helper()

	var a *net.TCPAddr
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return l.Addr().(*net.TCPAddr).Port, nil
		}
	}
	return
}

// WaitUntilBusyPort blocks until port is in use
func WaitUntilBusyPort(port string, t *testing.T) {
	t.Helper()
	startTime := time.Now()

	for {
		select {
		case <-time.After(100 * time.Millisecond):
			l, err := net.Listen("tcp", port)
			if err != nil {
				// Port is in use or unavailable
				if time.Since(startTime) > (100 * time.Millisecond) {
					// Timeout reached
					t.Logf("Server is listening on port %s", port)
					return
				}
				continue
			}
			l.Close()
			return
		}
	}
}
