package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/HamzaRahmani/urlShortner/internal/config"
	"github.com/HamzaRahmani/urlShortner/internal/database"
	"github.com/HamzaRahmani/urlShortner/internal/manager"
	"github.com/HamzaRahmani/urlShortner/internal/server"
)

func main() {
	fmt.Printf("Starting URL GO\n")
	srv, err := initializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app")
		return
	}

	err = srv.Start()
	if err != nil {
		log.Fatalf("Failed to start server")
		return
	}

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	log.Printf("Shutting Down")

	defer func() {
		srv.Stop()
	}()

}

func initializeApp() (*server.HTTPServer, error) {
	// "postgresql://postgres:gobank@localhost:5432/postgres"
	env := config.Init(os.Environ())
	port, _ := env.GetListeningPort()
	dbHost, _ := env.GetDatabaseHost()
	dbPass, _ := env.GetDatabasePassword()
	dbUser, _ := env.GetDatabaseUser()

	connString := fmt.Sprintf("postgresql://%s:%s@%s", dbUser, dbPass, dbHost)

	db, err := database.NewPostgresStore(connString)

	if err != nil {
		log.Printf("Failed to connect to db: %s", connString)
	}

	m := manager.NewManager(db)
	srv := server.NewHTTPServer(port, m)

	return srv, err
}
