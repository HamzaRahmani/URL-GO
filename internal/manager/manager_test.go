package manager_test

import (
	"regexp"
	"testing"

	"github.com/HamzaRahmani/urlShortner/internal/manager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateURL(t *testing.T) {
	t.Parallel()

	// Arrange
	inputURL := "https://www.google.ca/"
	database := new(mockDatabase)
	database.On("CreateURL", mock.MatchedBy(isURL)).Return(mock.MatchedBy(isURL), nil).Once()

	m := manager.NewManager(database)

	// Act
	hashedURL, err := m.CreateURL(inputURL)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedURL)
	assert.Less(t, hashedURL, inputURL)

	database.AssertExpectations(t)
}

func TestFindURL(t *testing.T) {

}

func isURL(input string) bool {
	// Define a regular expression for matching URLs
	urlRegex := regexp.MustCompile(`^(http|https):\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(\/\S*)?$`)

	// Use the MatchString function to check if the input matches the regex
	return urlRegex.MatchString(input)
}

type mockDatabase struct {
	mock.Mock
}
