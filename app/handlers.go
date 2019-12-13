package app

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func (s *server) roomsHandler() http.HandlerFunc {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		upgrader.Upgrade(w, r, nil)
	}
}
