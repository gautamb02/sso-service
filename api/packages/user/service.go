package user

import (
	"context"

	"github.com/gautamb02/sso-service/shared"
)

type UserService struct {
	userRepo UserRepositoryI
}

func NewUserService(repo UserRepositoryI) UserServiceI {
	return &UserService{
		userRepo: repo,
	}

}
func (s *UserService) RegisterUser(user *UserDetail, ctx context.Context) (int64, error) {
	exist, err := s.userRepo.CheckIfEmailExist(user.Email, ctx)
	if err != nil {
		return 0, err
	}
	if exist == 1 {
		return 0, shared.ErrEmailAlreadyExists
	}
	hashedPassword, err := shared.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashedPassword // Hash the password before saving
	return s.userRepo.RegisterUser(user, ctx)

}
