package service

import "github.com/expose443/social-network/internal/repository"

type ReactionServiceI interface {
}

type reactionService struct {
	reacRepo repository.ReactionRepositoryI
}
