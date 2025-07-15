package server

import (
	"fmt"

	"github.com/gautamb02/sso-service/api/packages/user"
	"github.com/gautamb02/sso-service/confreader"
	"github.com/gautamb02/sso-service/db"
	"github.com/gautamb02/sso-service/logger"
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
		user.NewUserModule(MongoDBs[shared.SsoService].DB),
	}

	router.SetupRoutes(handlers)
	// Start the server
	err := router.Run()
	if err != nil {
		logger.Warn("Failed to start server: %s", err.Error())
	}
}

func (s *Server) Setup() error {
	var err error

	MongoDBs[shared.SsoService], err = db.NewMongoClient(s.config.Databases.Mongos.SSO_Service)
	if err != nil {
		return fmt.Errorf("error establishing connection with %s: %s", shared.SsoService, err.Error())
	}

	return nil
}
