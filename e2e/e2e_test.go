package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type requestBody struct {
	URL    string `json:"url"`
	Expire bool   `json:"expiry"`
}

func TestCreateURL(t *testing.T) {
	body := requestBody{
		URL:    "https://www.google.ca/",
		Expire: false,
	}
	input, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/url", bytes.NewReader(input))
	resp := httptest.NewRecorder()

}

func TestGetURL(t *testing.T) {
	// Arrange

	// Act

	// Assert

}
