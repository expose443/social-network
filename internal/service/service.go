package service

import "github.com/expose443/social-network/internal/repository"

type ServiceI interface {
	NewAuthService(repo repository.AuthRepositoryI) AuthServiceI
}

type service struct {
	repo *repository.RepositoryI
}

func NewService(repo *repository.RepositoryI) ServiceI {
	return &service{
		repo: repo,
	}
}

func (s *service) NewAuthService(repo repository.AuthRepositoryI) AuthServiceI {
	return &authService{
		repo: repo,
	}
}
