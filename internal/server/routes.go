package server

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func NewRouter(db *pgx.Conn) *http.ServeMux {
	mux := http.DefaultServeMux
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	  fmt.Fprintf(w, "Hello, World!")
	})
	return mux
}
