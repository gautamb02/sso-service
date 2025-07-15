package server

import (
	"fmt"

	"github.com/gautamb02/sso-service/api/packages/user"
	"github.com/gautamb02/sso-service/confreader"
	"github.com/gautamb02/sso-service/db"
	"github.com/gautamb02/sso-service/rest"
	"github.com/gautamb02/sso-service/shared"
)

var MongoDBs = make(map[string]*db.MongoClient)

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
	handlers := []rest.IHTTPHandlerProvider{
		user.NewUserModule(MongoDBs[shared.SSO_SERVICE].DB),
	}

	router.SetupRoutes(handlers)
	// Start the server
	router.Run()
}

func (s *Server) Setup() error {
	var err error

	MongoDBs[shared.SSO_SERVICE], err = db.NewMongoClient(s.config.Databases.Mongos.SSO_Service)
	if err != nil {
		return fmt.Errorf("error establishing connection with %s: %s", shared.SSO_SERVICE, err.Error())
	}

	return nil
}
