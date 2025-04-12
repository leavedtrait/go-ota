package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/leavedtrait/go-ota/internal/views"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	// add logic to pass error messages
	component := views.LoginPage("")
	templ.Handler(component).ServeHTTP(w, r)
}
