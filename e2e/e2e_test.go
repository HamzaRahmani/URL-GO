package e2e

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"

	"github.com/HamzaRahmani/urlShortner/internal/config"
	"github.com/HamzaRahmani/urlShortner/internal/database"
	"github.com/HamzaRahmani/urlShortner/internal/manager"
	"github.com/HamzaRahmani/urlShortner/internal/server"
	"github.com/HamzaRahmani/urlShortner/internal/tests"
)

type requestBody struct {
	URL    string `json:"url"`
	Expire bool   `json:"expiry"`
}

func TestCreateURL(t *testing.T) {
	port, _ := tests.GetFreeTCPPort(t)

	db, err := database.NewPostgresStore(CreateConnString())
	if err != nil {
		panic(err)
	}
	m := manager.NewManager(db)
	srv := server.NewHTTPServer(port, m)
	srv.Start()
	defer srv.Stop()
	tests.WaitUntilBusyPort(port, t)

	body := requestBody{
		URL:    "https://www.google.ca/",
		Expire: false,
	}

	userRequest := httpexpect.Default(t, fmt.Sprintf("http://localhost:%d", port))

	userRequest.POST("/url").WithJSON(body).
		Expect().
		Status(http.StatusCreated).
		JSON().Object().ContainsKey("shortURL")
}

func CreateConnString() string {
	c := config.Init(os.Environ())
	user, _ := c.GetDatabaseUser()
	pass, _ := c.GetDatabasePassword()
	host, _ := c.GetDatabaseHost()

	connString := fmt.Sprintf("postgresql://%s:%s@%s?sslmode=disable", user, pass, host)
	return connString
}

// func TestGetURL(t *testing.T) {
// 	// Arrange

// 	// Act

// 	// Assert

// }
