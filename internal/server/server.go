package server

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"
)

// TODO: Inject next layer into HTTPServer
type HTTPServer struct {
	server *http.Server
}

// TODO: Create a handler
func NewHTTPServer(port int, router http.Handler) *HTTPServer {
	return &HTTPServer{
		&http.Server{
			Addr:              "localhost:" + strconv.Itoa(port),
			Handler:           router,
			ReadHeaderTimeout: 3 * time.Second,
		},
	}
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
