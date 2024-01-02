package manager_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/HamzaRahmani/urlShortner/internal/database"
	"github.com/HamzaRahmani/urlShortner/internal/manager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateURL(t *testing.T) {
	t.Parallel()
	inputURL := "https://www.google.ca"

	// Arrange
	db := new(mockDatabase)
	db.On(
		"InsertURL",
		mock.AnythingOfType("string"),
		mock.MatchedBy(isURL),
	).Return(database.URL{Hash: "abcdefg"}, nil).Once()

	m := manager.NewManager(db)

	// Act
	hashedURL, err := m.CreateURL(inputURL)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedURL)
	assert.Less(t, len(hashedURL), len(inputURL))
	assert.Len(t, getHash(hashedURL), 7)

	db.AssertExpectations(t)
}

func TestFindURL(t *testing.T) {

}

func TestCreateDuplicateURL(t *testing.T) {
	// Expected behaviour is to return the existing URL

}

func isURL(input string) bool {
	urlRegex := regexp.MustCompile(`^(http|https):\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(\/\S*)?$`)
	return urlRegex.MatchString(input)
}

func getHash(hashedURL string) string {
	lastSlashIndex := strings.LastIndex(hashedURL, "/")
	return hashedURL[lastSlashIndex+1:]
}

type mockDatabase struct {
	mock.Mock
}

// CreateURL implements database.Database.
func (m *mockDatabase) InsertURL(hash string, originalURL string) (database.URL, error) {
	args := m.Called(hash, originalURL)
	return args.Get(0).(database.URL), args.Error(1)
}
