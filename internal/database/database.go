package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	CreateURL(hash string, originalURL string) error
	// GetURL(hashedURL string) error
	// DeleteURL(hashedURL string) error
}

type PostgresStore struct {
	db *pgxpool.Pool
}

// TODO: define postgres connection info and inject via config layer
func NewPostgresStore(connString string) (*PostgresStore, error) {
	// os.Getenv("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	err = dbpool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStore{
		db: dbpool,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createURLTable()
}

func (s *PostgresStore) createURLTable() error {
	query := `create table if not exists url (
		hash char(7) primary key NOT NULL,
		original_url varchar(50) NOT NULL,
		created_at timestamp default current_timestamp
	)`

	_, err := s.db.Exec(context.TODO(), query)
	if err == nil {
		log.Printf("Successfully created table url")
	}

	return err
}

func (s PostgresStore) CreateURL(hash string, originalURL string) error {
	// query := `
	// insert into url
	// (raw_url, hashed_url)
	// values ($1, $2)
	// `

	return nil
}
func (s PostgresStore) DeleteURL(hashedURL string) error {
	return nil
}
func (s PostgresStore) GetURL(hashedURL string) (string, error) {
	return "", nil
}

// docker run --name some-postgres -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres
type URL struct {
	ID        int       `json:"id"`
	RawURL    string    `json:"rawURL"`
	HashedURL string    `json:"hashedURL"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewURL(rawURL, hashedURL string) *URL {
	return &URL{
		RawURL:    rawURL,
		HashedURL: hashedURL,
		CreatedAt: time.Now().UTC(),
	}
}
