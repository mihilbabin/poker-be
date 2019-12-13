package app

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mihilbabin/poker/config"
	"github.com/mihilbabin/poker/handlers"
)

type app struct {
	server *http.Server
	config *config.Config
	router chi.Router
}

func New() *app {
	conf := config.Init()
	router := chi.NewRouter()
	s := &http.Server{
		Addr:           conf.Server.Host,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        router,
	}
	a := &app{
		config: conf,
		router: router,
		server: s,
	}
	a.configureMiddlewares()
	a.configureRoutes()
	return a
}

func (a *app) Start() {
	log.Printf("Starting application server on %s", a.config.Server.Host)
	log.Fatal(a.server.ListenAndServe())
}

func (a *app) configureRoutes() {
	a.router.Get("/", handlers.Index)
	a.router.Get("/rooms", handlers.Rooms())
}

func (a *app) configureMiddlewares() {
	a.router.Use(middleware.RequestID)
	a.router.Use(middleware.RealIP)
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)
}
