package server

import (
	"context"
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
			Handler:           NewRouter(),
			ReadHeaderTimeout: 3 * time.Second,
		},
	}
}

// NewRouter routes all incoming requests.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(":)"))
	})

	return r
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
