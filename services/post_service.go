package services

import (
	"github.com/jinzhu/gorm"
	"webcms/models"
	"webcms/repositories"
)

type PostService interface {
	GetPostByID(id uint) (models.Post, error)
	UpdatePost(post models.Post) error
	DeletePost(id uint) error
	GetAllPosts() ([]models.Post, error)
}

type postService struct {
	postRepository repositories.PostRepository
	db             *gorm.DB
}

func NewPostService(postRepo repositories.PostRepository, db *gorm.DB) PostService {
	return &postService{postRepo, db}
}

func (s *postService) GetPostByID(id uint) (models.Post, error) {
	return s.postRepository.GetPostByID(id)
}

func (s *postService) UpdatePost(post models.Post) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.postRepository.UpdatePostWithTx(post, tx)
	})
}

func (s *postService) DeletePost(id uint) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.postRepository.DeletePostWithTx(id, tx)
	})
}

func (s *postService) GetAllPosts() ([]models.Post, error) {
	return s.postRepository.GetAllPosts()
}
