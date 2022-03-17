package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"pmain2/internal/config"
)

type Server struct {
	Host   string
	Port   string
	Router *mux.Router
}

func Create(c *config.Config) *Server {
	router := mux.NewRouter()
	return &Server{
		Host:   c.Host,
		Port:   c.Port,
		Router: router,
	}
}

func (s *Server) Run() error {
	err := http.ListenAndServe(":"+s.Port, s.Router)
	return err
}
