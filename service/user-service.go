package service

import "url-shortener/repository"

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}
