package services

import (
	"fmt"
	"github.com/Panda-ManR/article-backend/cmd/article_backend/models"
)

type postDAO interface {
	Get(id uint) (*models.Post, error)
	FindAll() []models.Post
}

type PostService struct {
	dao postDAO
}

// NewPostService creates a new PostService with the given post DAO.
func NewPostService(dao postDAO) *PostService {
	return &PostService{dao}
}

func (s *PostService) Get(id uint) (*models.Post, error) {
	return s.dao.Get(id)
}

func (s *PostService) FindAll() ([]models.Post, error) {
	posts := s.dao.FindAll()
	err := fmt.Errorf("no posts found")
	if len(posts) > 0 {
		return posts, nil
	}
	return posts, err
}
