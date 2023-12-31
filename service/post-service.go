package service

import (
	"github.com/kuma-coffee/go-search-from-postgresql/entity"
	"github.com/kuma-coffee/go-search-from-postgresql/repository"
)

type PostService interface {
	Store(post *entity.Post) error
	FindAll() ([]entity.Post, error)
	Search(query []string) ([]entity.Post, error)
	Sort(query string) ([]entity.Post, error)
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

func (s *service) FindAll() ([]entity.Post, error) {
	return s.postRepo.FindAll()
}

func (s *service) Search(query []string) ([]entity.Post, error) {
	return s.postRepo.Search(query)
}

func (s *service) Sort(query string) ([]entity.Post, error) {
	return s.postRepo.Sort(query)
}
