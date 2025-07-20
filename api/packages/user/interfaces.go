package user

import "context"

type UserServiceI interface {
	RegisterUser(user *UserDetail, ctx context.Context) (int64, error)
}

type UserRepositoryI interface {
	RegisterUser(user *UserDetail, ctx context.Context) (int64, error)
	CheckIfEmailExist(email string, ctx context.Context) (int64, error)
}
