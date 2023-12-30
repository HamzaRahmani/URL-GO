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

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(port int, manager manager.Manager) *HTTPServer {
	return &HTTPServer{
		&http.Server{
			Addr:              "localhost:" + strconv.Itoa(port),
			Handler:           NewRouter(),
			ReadHeaderTimeout: 3 * time.Second,
		},
	}
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(":)"))
	})

	return r
}

func (h *HTTPServer) Start() error {
	l, err := net.Listen("tcp4", h.server.Addr)
	if err != nil {
		return err
	}

	go func() { err = h.server.Serve(l) }()
	return err
}

func (h *HTTPServer) Stop() error {
	ctx := context.Background()
	err := h.server.Shutdown(ctx)
	return err
}
