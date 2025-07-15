package rest

import (
	"fmt"

	"github.com/gautamb02/sso-service/logger"
	"github.com/gin-gonic/gin"
)

type Route interface {
	SetupRoutes([]IHTTPHandlerProvider)
	Run()
}
type Router struct {
	Route
	rr *gin.Engine
}

func NewRouter() Route {
	rr := gin.Default()
	return &Router{
		rr: rr,
	}
}

func (r *Router) SetupRoutes(handlers []IHTTPHandlerProvider) {
	logger.Info("Setting up routes...")
	for _, handler := range handlers {
		for _, httpHandler := range handler.GetHTTPHandler() {
			path := fmt.Sprintf("/v%d/%s", httpHandler.Version, httpHandler.Path)
			logger.Info("Route: %s %s ", httpHandler.Method, path)
			r.rr.Handle(httpHandler.Method, path, APIWrapper(httpHandler.Func))
		}
	}
}
func (r *Router) Run() {
	r.rr.Run("0.0.0.0:1512")
}

func APIWrapper(function func(c *SessionContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionContext := NewSessionContext(c)
		function(sessionContext)
	}
}
