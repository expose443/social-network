package service

import "github.com/expose443/social-network/internal/repository"

type PostServiceI interface {
}

type postService struct {
	postRepo repository.PostRepositoryI
}
