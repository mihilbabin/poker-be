package main

import (
	"fmt"
	"net/http"

	"github.com/mihilbabin/poker/config"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Couldn't load .env file. Defaults will be used instead.")
	}
}

func main() {
	router := chi.NewRouter()
	c := config.Config()
	router.Get("/", hello)
	log.Printf("Starting application server on %s", c.Server.Host)
	log.Fatal(http.ListenAndServe(c.Server.Host, router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Chi")
}
