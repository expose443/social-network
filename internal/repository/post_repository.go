package repository

import "gorm.io/gorm"

type PostRepositoryI interface {
}

type postRepository struct {
	db *gorm.DB
}
