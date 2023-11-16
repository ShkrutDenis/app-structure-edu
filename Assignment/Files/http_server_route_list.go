package main

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) InitRoutes() {
	// list of endpoints
}

func (s *Server) Run() {}

func (s *Server) Shutdown() {}
