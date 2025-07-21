package user

import (
	"context"

	auth "github.com/gautamb02/sso-service/shared/authorization"
)

type UserServiceI interface {
	RegisterUser(user *auth.UserCreateRequest, ctx context.Context) (int64, error)
}

type UserRepositoryI interface {
	RegisterUser(user *auth.UserCreateRequest, ctx context.Context) (int64, error)
	CheckIfEmailExist(email string, ctx context.Context) (int64, error)
}
