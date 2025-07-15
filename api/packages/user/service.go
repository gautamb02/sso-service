package user

type UserService struct {
	userRepo UserRepositoryI
}

func NewUserService(repo UserRepositoryI) UserServiceI {
	return &UserService{
		userRepo: repo,
	}

}
func (s *UserService) RegisterUser() {

}
