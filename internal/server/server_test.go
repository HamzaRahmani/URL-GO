package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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

// TODO: Write test for handler, use chi and create an endpoint for create URL
// TODO: Fix this test, the chi router should be modular... maybe get familiar with chi and then progress with writing test
func TestCreateShortURL(t *testing.T) {
	// Arrange
	port, _ := tests.GetFreeTCPPort(t)
	handler = routeHandler{}

	srv := server.NewHTTPServer(port, handler)
	go func() { _ = srv.Start() }()
	defer func() { _ = srv.Stop() }()
	tests.WaitUntilBusyPort(string(rune(port)), t)

	json, _ := json.Marshal(requestBody{
		Page: "http://www.google.com",
	})
	body := bytes.NewReader(json)

	// Act
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("http://127.0.0.1:%d/short-url", port), body)
	w := httptest.NewRecorder()

	// res, err := http.Post(fmt.Sprintf("", port))
	defer func() { _ = res.Body.Close() }()

	// Assert
	assert.NoError(t, err)

}

type requestBody struct {
	Page string `json:"page"`
}
