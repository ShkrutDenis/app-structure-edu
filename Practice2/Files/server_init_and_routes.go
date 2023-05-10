package main

import "net/http"

// File that contains NewServer method for the creation HTTP server and also list of routes

type Server struct {
	Engine *http.Server
}

func NewServer(args ...interface{}) *Server {
	s := Server{
		Engine: &http.Server{},
	}
	s.InitRoutes()
	
	return &s
}

func (s *Server) InitRoutes() {
	// list of handlers
}

func (s *Server) Run() error {
	return s.Engine.ListenAndServe()
}
