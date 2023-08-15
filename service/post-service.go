package service

import (
	"github.com/kuma-coffee/go-search-from-postgresql/entity"
	"github.com/kuma-coffee/go-search-from-postgresql/repository"
)

type PostService interface {
	Store(post *entity.Post) error
}

type service struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return &service{postRepo}
}

func (s *service) Store(post *entity.Post) error {
	return s.postRepo.Store(post)
}
