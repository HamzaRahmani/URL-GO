package e2e

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"

	"github.com/HamzaRahmani/urlShortner/internal/server"
)

type requestBody struct {
	URL    string `json:"url"`
	Expire bool   `json:"expiry"`
}

func TestCreateURL(t *testing.T) {
	srv := server.NewHTTPServer()
	srv.Start()
	defer srv.Stop()

	body := requestBody{
		URL:    "https://www.google.ca/",
		Expire: false,
	}

	e := httpexpect.Default(t, "http://localhost:4000")

	e.POST("/url").WithJSON(body).
		Expect().
		Status(http.StatusCreated).
		JSON().Object().ContainsKey("shortenedURL")
}

// func TestGetURL(t *testing.T) {
// 	// Arrange

// 	// Act

// 	// Assert

// }
