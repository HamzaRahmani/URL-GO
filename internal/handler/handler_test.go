package handler_test

import (
	"testing"

	"github.com/HamzaRahmani/urlShortner/internal/handler"
	"github.com/stretchr/testify/mock"
)

func TestCreateURLHandler(t *testing.T) {
	// Arrange
	manager := new(mockManager)
	handler := handler.NewRouter()
}

type mockManager struct {
	mock.Mock
}
