package main

import (
	"fmt"
	"log"

	"github.com/HamzaRahmani/urlShortner/internal/database"
)

func main() {
	store, err := database.NewPostgresStore("postgres:gobank@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", store)
}
