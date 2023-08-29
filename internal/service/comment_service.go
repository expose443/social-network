package service

import "github.com/expose443/social-network/internal/repository"

type CommentServiceI interface{}

type commentService struct {
	commRepo repository.CommentRepositoryI
}
