package repository

import "gorm.io/gorm"

type RepositoryI interface {
	NewAuthRepository() AuthRepositoryI
	NewUserRepository() UserRepositoryI
	NewPostRepository() PostRepositoryI
	NewCommentRepository() CommentRepositoryI
	NewReactionRepository() ReactionRepositoryI
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) RepositoryI {
	return &repository{
		db: db,
	}
}

func (r *repository) NewAuthRepository() AuthRepositoryI {
	return &authRepository{
		db: r.db,
	}
}

func (r *repository) NewUserRepository() UserRepositoryI {
	return &userRepository{
		db: r.db,
	}
}

func (r *repository) NewPostRepository() PostRepositoryI {
	return &postRepository{
		db: r.db,
	}
}

func (r *repository) NewCommentRepository() CommentRepositoryI {
	return &commentRepository{
		db: r.db,
	}
}

func (r *repository) NewReactionRepository() ReactionRepositoryI {
	return &reactionRepository{
		db: r.db,
	}
}
