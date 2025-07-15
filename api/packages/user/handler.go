package user

import (
	"net/http"

	"github.com/gautamb02/sso-service/rest"
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
	err := c.BindJSON(&req)
	if err != nil {
		c.Respond(http.StatusBadRequest, map[string]string{"error": "error from bad request"})
		return
	}
	c.Respond(http.StatusOK, map[string]string{"hello": "from user package"})
}
