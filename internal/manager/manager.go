package manager

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"

	"github.com/HamzaRahmani/urlShortner/internal/database"
)

type Manager interface {
	CreateURL(rawURL string) (string, error)
	GetURL(hash string) (string, error)
}

type manager struct {
	database database.Database
	// TODO: define domain and inject via config layer
}

func NewManager(dB database.Database) *manager {
	return &manager{database: dB}
}

func (m *manager) CreateURL(originalURL string) (string, error) {
	md5 := getMD5Hash(originalURL)
	hash := encodeToBase62(md5)[:7]

	_, err := m.GetURL(hash)
	if err == nil {
		return hash, err
	}

	row, err := m.database.InsertURL(hash, originalURL)

	if err != nil {
		return "", err
	}

	return row.Hash, err
}
func (m *manager) GetURL(hash string) (string, error) {
	row, err := m.database.FindURL(hash)

	return row.OriginalURL, err
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encodeToBase62(text string) string {
	var i big.Int
	i.SetBytes([]byte(text))
	return i.Text(62)
}
