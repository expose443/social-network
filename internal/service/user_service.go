package service

import "github.com/expose443/social-network/internal/repository"

type UserServiceI interface {
}

type userService struct {
	userRepo repository.UserRepositoryI
}
