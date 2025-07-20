package user

import (
	"context"
	"net/http"

	"github.com/gautamb02/sso-service/logger"
	"github.com/gautamb02/sso-service/rest"
	"github.com/gautamb02/sso-service/shared"
)

type UserHandler struct {
	service UserServiceI
}

func NewUserHandler(s UserServiceI) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (uh *UserHandler) GetHTTPHandler() []*rest.HTTPHandler {
	return []*rest.HTTPHandler{
		{
			Version: 1,
			Method:  http.MethodGet,
			Path:    "api/users",
			Func:    uh.Homeuser,
		},
		{
			Version: 1,
			Method:  http.MethodPost,
			Path:    "api/signup",
			Func:    uh.Signup,
		},
	}
}

func (uh *UserHandler) Homeuser(c *rest.SessionContext) {
	c.Respond(http.StatusOK, map[string]string{"hello": "from user package"})
}

func (uh *UserHandler) Signup(c *rest.SessionContext) {
	var req UserDetail
	req.Verified = false // Default value for Verified field
	ctx := context.Background()
	// defer cancel()
	err := c.BindJSON(&req)
	if err != nil {
		c.Respond(http.StatusBadRequest, map[string]string{"error": "error from bad request"})
		return
	}
	_, err = uh.service.RegisterUser(&req, ctx)
	if err != nil {
		logger.Error("error while registering user %v", err.Error())
		if err == shared.ErrEmailAlreadyExists {
			c.Respond(http.StatusConflict, map[string]string{"error": "email already exists"})
			return
		}
		c.Respond(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}
	c.Respond(http.StatusCreated, map[string]string{"account": "created successfully"})
}
