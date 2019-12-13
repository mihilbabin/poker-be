package app

func (s *server) routes() {
	s.router.Get("/rooms", s.roomsHandler())
}
