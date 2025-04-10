package server

import (
	"log/slog"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(r.Method + " http://" + r.Host + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}