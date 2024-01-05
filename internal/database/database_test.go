package database_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/HamzaRahmani/urlShortner/internal/database"
	"github.com/HamzaRahmani/urlShortner/internal/tests"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

var testDbInstance *pgxpool.Pool
var testDbAddress string

func TestMain(m *testing.M) {
	testDB := tests.SetupTestDatabase()
	testDbInstance = testDB.DbInstance
	testDbAddress = testDB.DbAddress
	defer testDB.TearDown()
	os.Exit(m.Run())
}

func TestInsertURL(t *testing.T) {
	// Arrange
	db, err := database.NewPostgresStore(
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			tests.DbUser,
			tests.DbPass,
			testDbAddress,
			tests.DbName),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Act
	var row database.URL
	row, err = db.InsertURL("abcdefg", "http://longurl.ca")
	if err != nil {
		t.Errorf("error was not expected while updating: %s", err)
	}

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, row)
	assert.NotEmpty(t, row.CreatedAt)
}

func TestFindURL(t *testing.T) {
	// Arrange
	db, err := database.NewPostgresStore(
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			tests.DbUser,
			tests.DbPass,
			testDbAddress,
			tests.DbName),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Act
	var row database.URL
	row, err = db.FindURL("abcdefg")
	if err != nil {
		t.Errorf("error was not expected while updating: %s", err)
	}

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, row)
	assert.NotEmpty(t, row.OriginalURL)
	assert.NotEmpty(t, row.CreatedAt)
}
