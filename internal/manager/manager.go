package manager

import (
	"github.com/HamzaRahmani/urlShortner/internal/database"
)

type Manager interface {
	CreateURL(rawURL string) (string, error)
}

type manager struct {
	database database.Database
	// analyzer Analyzer
}

func NewManager(dB database.Database) *manager {
	return &manager{database: dB}
}

func (m *manager) CreateURL(text string) (string, error) {
	// hash url
	// insert hashed url into DB
	// return hashed url to user
	return "", nil
}
