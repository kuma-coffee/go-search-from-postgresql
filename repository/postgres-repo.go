package repository

import (
	"github.com/kuma-coffee/go-search-from-postgresql/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	Store(post *entity.Post) error
	FindAll() ([]entity.Post, error)
	Search(query []string) ([]entity.Post, error)
	Reset(table string) error
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

func (r *repo) Search(query []string) ([]entity.Post, error) {
	result := []entity.Post{}

	for _, val := range query {

		err := r.db.Table("posts").Where(
			r.db.Where("ndex LIKE ?", val).Or("pokemon LIKE ?", "%"+val+"%"),
		).Scan(&result).Error
		if err != nil {
			return nil, err
		}

		err = r.db.Table("posts").Where("? = ANY(type)", val).Scan(&result).Error
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (r *repo) Reset(table string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("TRUNCATE " + table).Error; err != nil {
			return err
		}

		return nil
	})
}
