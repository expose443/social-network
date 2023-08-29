package service

import "github.com/expose443/social-network/internal/repository"

type AuthServiceI interface {
	Register()
}

type authService struct {
	authRepo repository.AuthRepositoryI
}

func (a *authService) Register() {
	a.authRepo.Register()
}
