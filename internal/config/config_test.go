package config_test

import (
	"fmt"
	"testing"

	"github.com/HamzaRahmani/urlShortner/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestGetListeningPort(t *testing.T) {

	for _, tc := range []int{5125, 4024} {
		// Act
		c := config.Init([]string{fmt.Sprintf("LISTENING_PORT=%d", tc)})
		port, err := c.GetListeningPort()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, port, tc)
	}

}

func TestGetDatabaseHost(t *testing.T) {
	for _, expectedHost := range []string{"localhost:6379", "127.0.0.1:2750"} {
		expectedHost := expectedHost
		t.Run(fmt.Sprintf("with host %s", expectedHost), func(t *testing.T) {
			t.Parallel()

			// Act
			config := config.Init([]string{fmt.Sprintf("DATABASE_HOST=%s", expectedHost)})
			actualHost, err := config.GetDatabaseHost()

			// Assert
			assert.Equal(t, expectedHost, actualHost)
			assert.NoError(t, err)
		})
	}
}

func TestGetDatabaseUser(t *testing.T) {
	for _, expectedHost := range []string{"Danny", "Dani"} {
		expectedHost := expectedHost
		t.Run(fmt.Sprintf("with host %s", expectedHost), func(t *testing.T) {
			t.Parallel()

			// Act
			config := config.Init([]string{fmt.Sprintf("DATABASE_USER=%s", expectedHost)})
			actualHost, err := config.GetDatabaseUser()

			// Assert
			assert.Equal(t, expectedHost, actualHost)
			assert.NoError(t, err)
		})
	}
}

func TestGetDatabasePassword(t *testing.T) {
	for _, expectedPassword := range []string{"", "some_password"} {
		expectedPassword := expectedPassword
		t.Run(fmt.Sprintf("with host %s", expectedPassword), func(t *testing.T) {
			t.Parallel()

			// Act
			config := config.Init([]string{fmt.Sprintf("DATABASE_PASSWORD=%s", expectedPassword)})
			actualPassword, err := config.GetDatabasePassword()

			// Assert
			assert.Equal(t, expectedPassword, actualPassword)
			assert.NoError(t, err)
		})
	}
}
