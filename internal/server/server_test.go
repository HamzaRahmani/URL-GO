package server_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/HamzaRahmani/urlShortner/internal/server"
	"github.com/HamzaRahmani/urlShortner/internal/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServerStart(t *testing.T) {
	// Arrange
	port, _ := tests.GetFreeTCPPort(t)
	srv := server.NewHTTPServer(port, nil)
	go func() { _ = srv.Start() }()
	defer func() { _ = srv.Stop() }()
	tests.WaitUntilBusyPort(port, t)

	// Act
	res, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
	defer func() { _ = res.Body.Close() }()

	// Assert
	assert.NoError(t, err)
}

func TestStopServer(t *testing.T) {
	// Arrange
	port, _ := tests.GetFreeTCPPort(t)
	srv := server.NewHTTPServer(port, nil)
	go func() { _ = srv.Start() }()
	tests.WaitUntilBusyPort(port, t)

	// Act
	srv.Stop()
	_, err := http.Get(fmt.Sprintf("http://localhost:%d", int(port)))

	// Assert
	assert.NotNil(t, err)
}

func TestCreateURLHandler(t *testing.T) {
	// Arrange
	urlManager := new(mockManager)

	port, _ := tests.GetFreeTCPPort(t)
	srv := server.NewHTTPServer(port, urlManager)
	go func() { _ = srv.Start() }()
	tests.WaitUntilBusyPort(port, t)

	// Act
	srv.Stop()
	_, err := http.Post(fmt.Sprintf("http://localhost:%d/url", int(port)), "application/json", nil)

	// Assert
	assert.NotNil(t, err)
}

type mockManager struct {
	mock.Mock
}

func (m *mockManager) CreateURL(message string) (string, error) {
	args := m.Called(message)
	return args.Get(0).(string), args.Error(1)
}
