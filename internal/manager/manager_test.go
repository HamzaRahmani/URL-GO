package manager_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/HamzaRahmani/urlShortner/internal/manager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateURL(t *testing.T) {
	t.Parallel()
	inputURL := "https://www.google.ca/"

	// Arrange
	database := new(mockDatabase)
	database.On("CreateURL", mock.MatchedBy(isURL)).Return(mock.MatchedBy(isURL), nil).Once()

	m := manager.NewManager(database)

	// Act
	hashedURL, err := m.CreateURL(inputURL)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedURL)
	assert.Less(t, hashedURL, inputURL)
	assert.Len(t, getHash(hashedURL), 7)

	database.AssertExpectations(t)
}

func TestFindURL(t *testing.T) {

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
func (m *mockDatabase) CreateURL(url string) (string, error) {
	args := m.Called(url)
	return args.String(0), args.Error(1)
}
