package rest

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type SessionContext struct {
	c *gin.Context
}

func NewSessionContext(ctx *gin.Context) *SessionContext {
	return &SessionContext{
		c: ctx,
	}
}

func (sc *SessionContext) Respond(status int, data interface{}) {
	sc.c.JSON(status, data)
}

func (sc *SessionContext) BindJSON(req interface{}) error {
	return sc.c.ShouldBindJSON(req)
}

func (sc *SessionContext) Context() *gin.Context {
	return sc.c
}

func (sc *SessionContext) WithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(sc.c.Request.Context(), timeout)
}
