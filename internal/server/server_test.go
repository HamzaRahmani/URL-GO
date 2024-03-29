package server_test

import (
	"bytes"
	"encoding/json"
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

	urlManager.On("CreateURL", "https://www.google.ca/").Return("urlGO", nil).Once()

	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(struct {
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

	var body responseBody
	err = json.NewDecoder(resp.Body).Decode(&body)

	// Assert
	assert.Equal(t, resp.StatusCode, http.StatusCreated)
	assert.NoError(t, err, "Error decoding JSON")
	expectedBdoy := responseBody{Hash: "urlGO"}
	assert.Equal(t, expectedBdoy, body, "Unexpected response data")

	urlManager.AssertExpectations(t)
}

func TestGetURLHandler(t *testing.T) {
	// Arrange
	urlManager := new(mockManager)

	port, _ := tests.GetFreeTCPPort(t)
	srv := server.NewHTTPServer(port, urlManager)
	go func() { _ = srv.Start() }()
	defer func() { _ = srv.Stop() }()
	tests.WaitUntilBusyPort(port, t)

	expectedURL := "https://www.google.ca/"
	urlManager.On("GetURL", "xyzabcd").Return(expectedURL, nil).Once()

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/%s", port, "xyzabcd"))
	if err != nil {
		panic(err)
	}

	// Assert
	location := resp.Request.Response.Header.Get("Location")
	redirectStatusCode := resp.Request.Response.StatusCode
	destinationURL := resp.Request.URL

	assert.Equal(t, expectedURL, location, "Location header does not match expected value")
	assert.Equal(t, http.StatusMovedPermanently, redirectStatusCode)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, expectedURL, destinationURL.Host, "URL does not match expected value")

	urlManager.AssertExpectations(t)
}

func TestURLValidator(t *testing.T) {
	// Arrange
	urlManager := new(mockManager)

	port, _ := tests.GetFreeTCPPort(t)
	srv := server.NewHTTPServer(port, urlManager)
	go func() { _ = srv.Start() }()
	defer func() { _ = srv.Stop() }()
	tests.WaitUntilBusyPort(port, t)

	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(struct {
		URL string `json:"url"`
	}{
		"https://www.googl`/.ca/",
	})

	// Act
	resp, err := http.Post(fmt.Sprintf("http://localhost:%d/url", port), "application/json", &buf)

	if err != nil {
		panic(err)
	}

	// Assert
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
	urlManager.AssertExpectations(t)
}

type mockManager struct {
	mock.Mock
}

func (m *mockManager) CreateURL(rawURL string) (string, error) {
	args := m.Called(rawURL)
	return args.String(0), args.Error(1)
}

func (m *mockManager) GetURL(hash string) (string, error) {
	args := m.Called(hash)
	return args.String(0), args.Error(1)
}

// ResponseBody represents the structure of the expected JSON response.
type responseBody struct {
	Hash string `json:"hash"`
}
