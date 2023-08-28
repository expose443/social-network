package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type AuthRepositoryI interface {
	Register()
}

type authRepository struct {
	db *gorm.DB
}

func (a *authRepository) Register() {
	fmt.Println("test")

}
