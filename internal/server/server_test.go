package server_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/HamzaRahmani/urlShortner/internal/server"
	"github.com/HamzaRahmani/urlShortner/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestServerStart(t *testing.T) {
	// Arrange
	port, _ := tests.GetFreeTCPPort(t)
	srv := server.NewHTTPServer(port)
	go func() { _ = srv.Start() }()
	defer func() { _ = srv.Stop() }()
	tests.WaitUntilBusyPort(string(rune(port)), t)

	// Act
	res, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d", port))
	defer func() { _ = res.Body.Close() }()

	// Assert
	assert.NoError(t, err)
}

func TestStopServer(t *testing.T) {
	// Arrange
	port, _ := tests.GetFreeTCPPort(t)
	srv := server.NewHTTPServer(port)
	go func() { _ = srv.Start() }()
	tests.WaitUntilBusyPort(string(rune(port)), t)

	// Act
	srv.Stop()
	_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d", port))

	// Assert
	assert.NotNil(t, err)
}
