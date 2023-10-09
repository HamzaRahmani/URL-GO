package server_test

import (
	"net/http"
	"testing"

	"github.com/HamzaRahmani/urlShortner/internal/server"
	"github.com/gavv/httpexpect/v2"
)

// TODO: Test Server Setup, should be able to find an open port and spin up server, create a function that looks for an open port
// TODO: Test Server Teardown, after teardown, should not be pingable

func TestServerStart(t *testing.T) {
	srv := server.NewHTTPServer()
	go srv.Start()
	defer srv.Stop()

	userRequest := httpexpect.Default(t, "http://localhost:4000")

	userRequest.GET("/health-check").Expect().Status(http.StatusOK)

}

func TestStopServer(t *testing.T) {

}
