package e2e

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"

	"github.com/HamzaRahmani/urlShortner/internal/server"
	"github.com/HamzaRahmani/urlShortner/internal/tests"
)

type requestBody struct {
	URL    string `json:"url"`
	Expire bool   `json:"expiry"`
}

func TestCreateURL(t *testing.T) {
	port, _ := tests.GetFreeTCPPort(t)
	srv := server.NewHTTPServer(port)
	srv.Start()
	defer srv.Stop()

	body := requestBody{
		URL:    "https://www.google.ca/",
		Expire: false,
	}

	userRequest := httpexpect.Default(t, "http://localhost:4000")

	userRequest.POST("/url").WithJSON(body).
		Expect().
		Status(http.StatusCreated).
		JSON().Object().ContainsKey("shortenedURL")
}

// func TestGetURL(t *testing.T) {
// 	// Arrange

// 	// Act

// 	// Assert

// }
