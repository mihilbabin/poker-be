package config

import (
	"sync"
)

type Config struct {
	Server *serverConfig
}

type serverConfig struct {
	Host string
}

var conf *Config
var once sync.Once

func Init() *Config {
	once.Do(func() {
		conf = newConfig()
	})
	return conf
}

func newConfig() *Config {
	return &Config{
		Server: &serverConfig{
			Host: GetEnv("HOST").Str("localhost:8080"),
		},
	}
}
