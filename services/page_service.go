package services

import (
	"github.com/jinzhu/gorm"
	"webcms/models"
	"webcms/repositories"
)

type PageService interface {
	GetPageByID(id uint) (models.Page, error)
	UpdatePage(page models.Page) error
	DeletePage(id uint) error
	GetAllPages() ([]models.Page, error)
}

type pageService struct {
	pageRepository repositories.PageRepository
	db             *gorm.DB
}

func NewPageService(pageRepo repositories.PageRepository, db *gorm.DB) PageService {
	return &pageService{pageRepo, db}
}

func (s *pageService) GetPageByID(id uint) (models.Page, error) {
	return s.pageRepository.GetPageByID(id)
}

func (s *pageService) UpdatePage(page models.Page) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.pageRepository.UpdatePageWithTx(page, tx)
	})
}

func (s *pageService) DeletePage(id uint) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.pageRepository.DeletePageWithTx(id, tx)
	})
}

func (s *pageService) GetAllPages() ([]models.Page, error) {
	return s.pageRepository.GetAllPages()
}
