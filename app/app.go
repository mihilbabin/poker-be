package app

import (
	"net/http"

	"github.com/go-chi/chi"
)

type server struct {
	server *http.Server
	router chi.Router
}
