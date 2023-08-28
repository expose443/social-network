package repository

import "gorm.io/gorm"

type UserRepositoryI interface {
}

type userRepository struct {
	db *gorm.DB
}
