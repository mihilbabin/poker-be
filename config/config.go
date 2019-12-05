package config

import (
	"sync"
)

type config struct {
	Server *serverConfig
}

type serverConfig struct {
	Host string
}

var conf *config
var once sync.Once

func Config() *config {
	once.Do(func() {
		conf = newConfig()
	})
	return conf
}

func newConfig() *config {
	return &config{
		Server: &serverConfig{
			Host: GetEnv("HOST").Str("localhost:8080"),
		},
	}
}
