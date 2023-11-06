package server

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// TODO: Inject next layer into HTTPServer
type HTTPServer struct {
	server *http.Server
}

// TODO: Create a handler
func NewHTTPServer(port int) *HTTPServer {
	return &HTTPServer{
		&http.Server{
			Addr:              "0.0.0.0:" + strconv.Itoa(port),
			Handler:           Routes(),
			ReadHeaderTimeout: 3 * time.Second,
		},
	}
}

func Routes() http.Handler {
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
