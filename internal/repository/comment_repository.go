package repository

import "gorm.io/gorm"

type CommentRepositoryI interface {
}

type commentRepository struct {
	db *gorm.DB
}
