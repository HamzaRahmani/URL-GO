package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/HamzaRahmani/urlShortner/internal/manager"
	"github.com/go-chi/chi/v5"
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

	// r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("hello"))
	// })

	// r.Post("/url", func(w http.ResponseWriter, r *http.Request) {
	// 	var body createURLRequest
	// 	err := json.NewDecoder(r.Body).Decode(&body)
	// 	if err != nil {
	// 		http.Error(w, fmt.Sprintf(http.StatusText(400), ": ", err), 400)
	// 		return
	// 	}

	// 	shortURL, _ := m.CreateURL(body.URL)

	// 	data := &responseBody{ShortURL: shortURL}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusCreated)
	// 	json.NewEncoder(w).Encode(data)

	// })

	r.Get("/{hash}", func(w http.ResponseWriter, r *http.Request) {
		hash := chi.URLParam(r, "hash")

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

type createURLRequest struct {
	URL string `json:"url"`
}

type responseBody struct {
	ShortURL string `json:"shortURL"`
}

// Start starts the HTTP server.
func (h *HTTPServer) Start() error {
	l, err := net.Listen("tcp4", h.server.Addr)
	if err != nil {
		return err
	}

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
