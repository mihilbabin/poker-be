package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Couldn't load .env file. Defaults will be used instead.")
	}
}

func main() {
	s := newServer()
	s.Start()
}
