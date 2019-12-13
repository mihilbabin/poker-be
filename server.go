package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mihilbabin/poker/config"
	"github.com/mihilbabin/poker/handlers"
)

type server struct {
	c      *config.Config
	router chi.Router
}

func newServer() *server {
	s := &server{
		c:      config.Init(),
		router: chi.NewRouter(),
	}
	s.configureMiddlewares()
	s.configureRoutes()
	return s
}

func (s *server) Start() {
	log.Printf("Starting application server on %s", s.c.Server.Host)
	log.Fatal(http.ListenAndServe(s.c.Server.Host, s.router))
}

func (s *server) configureRoutes() {
	s.router.Get("/", handlers.Index)
}

func (s *server) configureMiddlewares() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
}
