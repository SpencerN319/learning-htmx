package server

func (s *server) routes() {
	s.router.GET("/health", s.handleHealth())
	s.router.GET("/", s.handleIndex())
}
