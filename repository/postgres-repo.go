package repository

import (
	"github.com/kuma-coffee/go-search-from-postgresql/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	Store(post *entity.Post) error
	FindAll() ([]entity.Post, error)
}

type repo struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) PostRepository {
	return &repo{db}
}

func (r *repo) Store(post *entity.Post) error {
	err := r.db.Create(&post).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	result := []entity.Post{}

	err := r.db.Table("posts").Select("*").Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
