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
	InsertURL(hash string, originalURL string) (URL, error)
	FindURL(hash string) (URL, error)
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
	}, err
}

func (s *PostgresStore) Init() error {
	return s.createURLTable()
}

func (s *PostgresStore) createURLTable() error {
	query := `create table if not exists url (
		hash char(7) primary key NOT NULL,
		original_url varchar(255) NOT NULL,
		created_at timestamp default current_timestamp
	)`

	_, err := s.db.Exec(context.TODO(), query)
	if err == nil {
		log.Printf("Successfully created table url")
	}

	return err
}

func (s PostgresStore) InsertURL(hash string, originalURL string) (URL, error) {
	query := `
	insert into url
	(hash, original_url)
	values ($1, $2)
	returning hash, original_url, created_at
	`
	var row URL
	err := s.db.QueryRow(context.Background(), query, hash, originalURL).
		Scan(&row.Hash, &row.OriginalURL, &row.CreatedAt)
	if err != nil {
		log.Fatalf("Failed to insert URL: %s", err)

	}

	return row, err
}

func (s PostgresStore) FindURL(hash string) (URL, error) {
	query := `
	select 	hash, original_url, created_at
	from url
	where hash = $1
	`
	var row URL
	err := s.db.QueryRow(context.Background(), query, hash).
		Scan(&row.Hash, &row.OriginalURL, &row.CreatedAt)
	if err != nil {
		log.Fatalf("Failed to find hash: %s", err)
	}
	return row, err
}

func (s PostgresStore) DeleteURL(hashedURL string) error {
	return nil
}

// docker run --name some-postgres -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres
type URL struct {
	Hash        string    `json:"hash"`
	OriginalURL string    `json:"hashedURL"`
	CreatedAt   time.Time `json:"createdAt"`
}
