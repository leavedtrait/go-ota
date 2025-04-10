package server

import (
	"net/http"
	"github.com/jackc/pgx/v5"
)

func NewRouter(db *pgx.Conn) *http.ServeMux {
	mux := http.DefaultServeMux
	

	return mux
}
