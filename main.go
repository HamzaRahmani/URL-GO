package main

import (
	"fmt"
	"log"

	"github.com/HamzaRahmani/urlShortner/internal/database"
)

func main() {
	// TODO: create a bootstrap function that starts the server, pass config into it from main
	store, err := database.NewPostgresStore("postgresql://postgres:gobank@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", store)
}
