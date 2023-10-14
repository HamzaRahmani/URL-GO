package server_test

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/HamzaRahmani/urlShortner/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestServerStart(t *testing.T) {
	// Arrange
	port, _ := GetFreeTCPPort(t)
	srv := server.NewHTTPServer(port)
	go func() { _ = srv.Start() }()
	defer func() { _ = srv.Stop() }()
	WaitUntilBusyPort(string(rune(port)), t)

	// Act
	res, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d", port))
	defer func() { _ = res.Body.Close() }()

	// Assert
	assert.NoError(t, err)
}

func TestStopServer(t *testing.T) {
	// Arrange
	port, _ := GetFreeTCPPort(t)
	srv := server.NewHTTPServer(port)
	go func() { _ = srv.Start() }()
	WaitUntilBusyPort(string(rune(port)), t)

	// Act
	srv.Stop()
	_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d", port))

	// Assert
	assert.NotNil(t, err)
}

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
