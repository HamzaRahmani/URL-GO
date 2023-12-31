package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

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
	_, err := http.Get(fmt.Sprintf("http://localhost:%d", port))

	// Assert
	assert.NotNil(t, err)
}

func TestCreateURLHandler(t *testing.T) {
	// Arrange
	urlManager := new(mockManager)

	port, _ := tests.GetFreeTCPPort(t)
	srv := server.NewHTTPServer(port, urlManager)
	go func() { _ = srv.Start() }()
	defer func() { _ = srv.Stop() }()
	tests.WaitUntilBusyPort(port, t)
	time.Sleep(10 * time.Millisecond)

	urlManager.On("CreateURL", "https://www.google.ca/").Return("urlGO", nil).Once()

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(struct {
		URL string `json:"url"`
	}{
		"https://www.google.ca/",
	})

	// Act
	resp, err := http.Post(fmt.Sprintf("http://localhost:%d/url", port), "application/json", &buf)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Assert
	assert.Equal(t, resp.StatusCode, http.StatusCreated)

	var body responseBody
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err, "Error decoding JSON")

	expectedBdoy := responseBody{ShortURL: "urlGO"}
	assert.Equal(t, expectedBdoy, body, "Unexpected response data")

	urlManager.AssertExpectations(t)
}

type mockManager struct {
	mock.Mock
}

func (m *mockManager) CreateURL(message string) (string, error) {
	args := m.Called(message)
	return args.String(0), args.Error(1)
}

// ResponseBody represents the structure of the expected JSON response.
type responseBody struct {
	ShortURL string `json:"shortURL"`
}
