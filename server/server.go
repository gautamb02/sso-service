package server

import (
	"net/http"

	"github.com/gautamb02/sso-service/confreader"
	"github.com/gautamb02/sso-service/rest"
)

type Server struct {
	config *confreader.Config
}

func NewServer(config *confreader.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start() {
	router := rest.NewRouter()
	router.R.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the SSO Service!"))
	})
	http.ListenAndServe(":3333", router.R)
}
