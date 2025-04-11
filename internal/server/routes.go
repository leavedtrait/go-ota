package server

import (
	"fmt"
	"net/http"

	"github.com/leavedtrait/go-ota/internal/db"
	"github.com/leavedtrait/go-ota/internal/handlers"
)

func NewRouter(queries *db.Queries) *http.ServeMux {
	mux := http.DefaultServeMux
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	  fmt.Fprintf(w, "Hello, World!")
	})
	mux.HandleFunc("POST /users/create", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUserHandler(queries, w, r)
	})
	mux.HandleFunc("PUT /users/update", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUserHandler(queries, w, r)
	})
	mux.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUserByIDHandler(queries, w, r)
	})
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUserByIDHandler(queries, w, r)
	})
	
	return mux
}
