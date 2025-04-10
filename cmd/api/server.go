package api

import (
	"log/slog"
	"net/http"
	"github.com/leavedtrait/go-ota/internal/server"
)

type Server struct {
	Srv http.Server
}

func NewServer(addr string, mux *http.ServeMux) *Server {
	handler := server.LoggingMiddleware(mux)
	return &Server{
		Srv: http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *Server) ListenAndServe() error {
	slog.Info("Server listening on http://localhost" + s.Srv.Addr)
	return s.Srv.ListenAndServe()
}
