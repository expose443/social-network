package repository

import "gorm.io/gorm"

type AuthRepositoryI interface {
	Register()
}

type authRepository struct {
	db *gorm.DB
}

func (a *authRepository) Register() {

}
