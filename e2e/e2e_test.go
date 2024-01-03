package e2e

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/HamzaRahmani/urlShortner/internal/database"
	"github.com/HamzaRahmani/urlShortner/internal/manager"
	"github.com/HamzaRahmani/urlShortner/internal/server"
	"github.com/HamzaRahmani/urlShortner/internal/tests"
)

type requestBody struct {
	URL    string `json:"url"`
	Expire bool   `json:"expiry"`
}

var testDbInstance *pgxpool.Pool
var testDbAddress string

func TestMain(m *testing.M) {
	testDB := tests.SetupTestDatabase()
	testDbInstance = testDB.DbInstance
	testDbAddress = testDB.DbAddress
	defer testDB.TearDown()
	os.Exit(m.Run())
}

func TestCreateURL(t *testing.T) {
	port, _ := tests.GetFreeTCPPort(t)
	// TODO: use ephemeral database used in database_test.go
	db, err := database.NewPostgresStore(
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			tests.DbUser,
			tests.DbPass,
			testDbAddress,
			tests.DbName),
	)

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

// func CreateConnString() string {
// 	c := config.Init(os.Environ())
// 	user, _ := c.GetDatabaseUser()
// 	pass, _ := c.GetDatabasePassword()
// 	host, _ := c.GetDatabaseHost()

// 	connString := fmt.Sprintf("postgresql://%s:%s@%s?sslmode=disable", user, pass, host)
// 	return connString
// }

// func TestGetURL(t *testing.T) {
// 	// Arrange

// 	// Act

// 	// Assert

// }
