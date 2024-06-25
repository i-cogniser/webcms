// services/content_service.go
package services

import (
	"webcms/models"
	"webcms/repositories"
)

type ContentService interface {
	CreatePost(post models.Post) error
	GetPostByID(id uint) (models.Post, error)
	UpdatePost(post models.Post) error
	DeletePost(id uint) error
	GetAllPosts() ([]models.Post, error)

	CreatePage(page models.Page) error
	GetPageByID(id uint) (models.Page, error)
	UpdatePage(page models.Page) error
	DeletePage(id uint) error
	GetAllPages() ([]models.Page, error)
}

type contentService struct {
	postRepository repositories.PostRepository
	pageRepository repositories.PageRepository
}

func NewContentService(postRepo repositories.PostRepository, pageRepo repositories.PageRepository) ContentService {
	return &contentService{postRepo, pageRepo}
}

func (s *contentService) CreatePost(post models.Post) error {
	return s.postRepository.CreatePost(post)
}

func (s *contentService) GetPostByID(id uint) (models.Post, error) {
	return s.postRepository.GetPostByID(id)
}

func (s *contentService) UpdatePost(post models.Post) error {
	return s.postRepository.UpdatePost(post)
}

func (s *contentService) DeletePost(id uint) error {
	return s.postRepository.DeletePost(id)
}

func (s *contentService) GetAllPosts() ([]models.Post, error) {
	return s.postRepository.GetAllPosts()
}

func (s *contentService) CreatePage(page models.Page) error {
	return s.pageRepository.CreatePage(page)
}

func (s *contentService) GetPageByID(id uint) (models.Page, error) {
	return s.pageRepository.GetPageByID(id)
}

func (s *contentService) UpdatePage(page models.Page) error {
	return s.pageRepository.UpdatePage(page)
}

func (s *contentService) DeletePage(id uint) error {
	return s.pageRepository.DeletePage(id)
}

func (s *contentService) GetAllPages() ([]models.Page, error) {
	return s.pageRepository.GetAllPages()
}
