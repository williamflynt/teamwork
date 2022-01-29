package server

import (
	"teamwork/internal/teamwork"
	"net/http"
)

type Server struct {
		Teamwork teamwork.App
}

func New(tw teamwork.App) (*Server, error) {
	return &Server{Teamwork: tw}, nil
}

func (s *Server) Run() error {
	http.Handle("/", http.HandlerFunc(healthcheckHandler()))
	return http.ListenAndServe(":8080",nil)
}