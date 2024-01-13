package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/HamzaRahmani/urlShortner/internal/manager"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// HTTPServer represents a new HTTP server
type HTTPServer struct {
	server *http.Server
}

// NewHTTPServer creates a new HTTP server configured with the provided port and manager.
func NewHTTPServer(port int, manager manager.Manager) *HTTPServer {
	return &HTTPServer{
		&http.Server{
			Addr:              "localhost:" + strconv.Itoa(port),
			Handler:           NewRouter(manager),
			ReadHeaderTimeout: 3 * time.Second,
		},
	}
}

// NewRouter routes all incoming requests.
func NewRouter(m manager.Manager) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})

	r.Post("/url", func(w http.ResponseWriter, r *http.Request) {
		var body createURLRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, fmt.Sprintf(http.StatusText(400), ": ", err), 400)
			return
		}

		if !isURL(body.URL) {
			http.Error(w, fmt.Sprintf("%s: input was not a URL", http.StatusText(400)), 400)
			return
		}

		hash, err := m.CreateURL(body.URL)

		if err != nil {
			http.Error(w, fmt.Sprintf("%s: failed to shorten URL", http.StatusText(500)), 500)
			return
		}

		data := &responseBody{Hash: hash}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)

	})

	r.Get("/{hash}", func(w http.ResponseWriter, r *http.Request) {
		hash := chi.URLParam(r, "hash")

		if !isValidLength(hash) {
			http.Error(w, fmt.Sprintf("%s: not a valid hash", http.StatusText(400)), 400)
			return
		}

		originalURL, err := m.GetURL(hash)
		if err != nil {
			http.Error(w, fmt.Sprintf(http.StatusText(404), ": ", err), 404)
			return
		}

		w.Header().Add("Location", originalURL)
		w.WriteHeader(http.StatusMovedPermanently)
	})

	return r
}

func isURL(input string) bool {
	urlRegex := regexp.MustCompile(`^(http|https):\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(\/\S*)?$`)
	return urlRegex.MatchString(input)
}

func isValidLength(input string) bool {
	return len(input) == 7
}

type createURLRequest struct {
	URL string `json:"url"`
}

type responseBody struct {
	Hash string `json:"hash"`
}

// Start starts the HTTP server.
func (h *HTTPServer) Start() error {
	l, err := net.Listen("tcp4", h.server.Addr)
	if err != nil {
		panic(err)
	}
	log.Printf("Listening at: %s", h.server.Addr)
	go func() { err = h.server.Serve(l) }()
	return err
}

// Stop gracefully shuts down the HTTP server by initiating a shutdown process
// and waiting for existing connections to complete.
func (h *HTTPServer) Stop() error {
	ctx := context.Background()
	err := h.server.Shutdown(ctx)
	return err
}
