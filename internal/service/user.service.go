package service

import "github.com/hungntth/be-go-ecom/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}