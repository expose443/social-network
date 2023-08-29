package service

import "github.com/expose443/social-network/internal/repository"

type ServiceI interface {
	NewCommentService(repo *repository.RepositoryI) CommentServiceI
	NewAuthService(repo *repository.RepositoryI) AuthServiceI
	NewReactionService(repo *repository.RepositoryI) ReactionServiceI
}

type service struct {
	repo *repository.RepositoryI
}

func NewService(repo *repository.RepositoryI) ServiceI {
	return &service{
		repo: repo,
	}
}

func (s *service) NewAuthService(repo *repository.RepositoryI) AuthServiceI {
	return &authService{
		authRepo: (*repo).NewAuthRepository(),
	}
}

func (s *service) NewCommentService(repo *repository.RepositoryI) CommentServiceI {
	return &commentService{
		commRepo: (*repo).NewCommentRepository(),
	}
}

func (s *service) NewUserService(repo *repository.RepositoryI) UserServiceI {
	return &userService{
		userRepo: (*repo).NewUserRepository(),
	}
}

func (s *service) NewPostService(repo *repository.RepositoryI) PostServiceI {
	return &postService{
		postRepo: (*repo).NewPostRepository(),
	}
}

func (s *service) NewReactionService(repo *repository.RepositoryI) ReactionServiceI {
	return &reactionService{
		reacRepo: (*repo).NewReactionRepository(),
	}
}
