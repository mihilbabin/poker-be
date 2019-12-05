package config

import (
	"os"
	"strconv"
)

type env struct {
	value  string
	exists bool
}

func GetEnv(name string) *env {
	value, exists := os.LookupEnv(name)
	return &env{value: value, exists: exists}
}

func (e *env) Str(fallback string) string {
	if e.exists {
		return e.value
	}
	return fallback
}

func (e *env) Int(fallback int) int {
	if e.exists {
		if value, err := strconv.Atoi(e.value); err == nil {
			return value
		}
	}
	return fallback
}

func (e *env) Bool(fallback bool) bool {
	if e.exists {
		if value, err := strconv.ParseBool(e.value); err == nil {
			return value
		}
	}
	return fallback
}
