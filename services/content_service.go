package services

import (
	"github.com/jinzhu/gorm"
	"webcms/models"
	"webcms/repositories"
)

type ContentService interface {
	CreatePost(post models.Post) error
	CreatePostWithTx(post models.Post, tx *gorm.DB) error
	GetPostByID(id uint) (models.Post, error)
	UpdatePost(post models.Post) error
	DeletePost(id uint) error
	GetAllPosts() ([]models.Post, error)
	CreatePage(page models.Page) error
	CreatePageWithTx(page models.Page, tx *gorm.DB) error
	GetPageByID(id uint) (models.Page, error)
	UpdatePage(page models.Page) error
	DeletePage(id uint) error
	GetAllPages() ([]models.Page, error)
}

type contentService struct {
	postRepository repositories.PostRepository
	pageRepository repositories.PageRepository
	db             *gorm.DB
}

func NewContentService(postRepo repositories.PostRepository, pageRepo repositories.PageRepository, db *gorm.DB) ContentService {
	return &contentService{postRepo, pageRepo, db}
}

func (s *contentService) CreatePost(post models.Post) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.CreatePostWithTx(post, tx)
	})
}

func (s *contentService) CreatePostWithTx(post models.Post, tx *gorm.DB) error {
	return s.postRepository.CreatePostWithTx(post, tx)
}

func (s *contentService) GetPostByID(id uint) (models.Post, error) {
	return s.postRepository.GetPostByID(id)
}

func (s *contentService) UpdatePost(post models.Post) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.UpdatePostWithTx(post, tx)
	})
}

func (s *contentService) UpdatePostWithTx(post models.Post, tx *gorm.DB) error {
	return s.postRepository.UpdatePostWithTx(post, tx)
}

func (s *contentService) DeletePost(id uint) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.DeletePostWithTx(id, tx)
	})
}

func (s *contentService) DeletePostWithTx(id uint, tx *gorm.DB) error {
	return s.postRepository.DeletePostWithTx(id, tx)
}

func (s *contentService) GetAllPosts() ([]models.Post, error) {
	return s.postRepository.GetAllPosts()
}

func (s *contentService) CreatePage(page models.Page) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.CreatePageWithTx(page, tx)
	})
}

func (s *contentService) CreatePageWithTx(page models.Page, tx *gorm.DB) error {
	return s.pageRepository.CreatePageWithTx(page, tx)
}

func (s *contentService) GetPageByID(id uint) (models.Page, error) {
	return s.pageRepository.GetPageByID(id)
}

func (s *contentService) UpdatePage(page models.Page) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.UpdatePageWithTx(page, tx)
	})
}

func (s *contentService) UpdatePageWithTx(page models.Page, tx *gorm.DB) error {
	return s.pageRepository.UpdatePageWithTx(page, tx)
}

func (s *contentService) DeletePage(id uint) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.DeletePageWithTx(id, tx)
	})
}

func (s *contentService) DeletePageWithTx(id uint, tx *gorm.DB) error {
	return s.pageRepository.DeletePageWithTx(id, tx)
}

func (s *contentService) GetAllPages() ([]models.Page, error) {
	return s.pageRepository.GetAllPages()
}
