package main

import (
	"github.com/joho/godotenv"
	"github.com/mihilbabin/poker/app"
	log "github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Couldn't load .env file. Defaults will be used instead.")
	}
}

func main() {
	a := app.New()
	a.Start()
}
