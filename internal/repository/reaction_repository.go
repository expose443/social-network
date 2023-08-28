package repository

import "gorm.io/gorm"

type ReactionRepositoryI interface {
}

type reactionRepository struct {
	db *gorm.DB
}
