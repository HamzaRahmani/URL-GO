package database_test

import (
	"os"
	"testing"

	"github.com/HamzaRahmani/urlShortner/internal/tests"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testDbInstance *pgxpool.Pool

func TestMain(m *testing.M) {
	testDB := tests.SetupTestDatabase()
	testDbInstance = testDB.DbInstance
	defer testDB.TearDown()
	os.Exit(m.Run())
}

// func TestCreateURL(t *testing.T) {
// 	// Arrange
// 	db, err := database.NewPostgresStore()
// 	mock.Conn().PgConn()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	mock.ExpectExec("INSERT INTO url").WithArgs("abcdefg", "http://longurl.ca")

// 	// Act
// 	if err = db.CreateURL("abcdefg", "http://longurl.ca"); err != nil {
// 		t.Errorf("error was not expected while updating: %s", err)
// 	}

// 	// Assert
// 	assert.NoError(t, err)
// }
