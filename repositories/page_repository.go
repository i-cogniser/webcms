package repositories

import (
	"webcms/models"

	"github.com/jinzhu/gorm"
)

type PageRepository interface {
	CreatePage(page models.Page) error
	CreatePageWithTx(page models.Page, tx *gorm.DB) error
	GetPageByID(id uint) (models.Page, error)
	UpdatePage(page models.Page) error
	UpdatePageWithTx(page models.Page, tx *gorm.DB) error
	DeletePage(id uint) error
	DeletePageWithTx(id uint, tx *gorm.DB) error
	GetAllPages() ([]models.Page, error)
}

type pageRepository struct {
	db *gorm.DB
}

func NewPageRepository(db *gorm.DB) PageRepository {
	return &pageRepository{db}
}

func (r *pageRepository) CreatePage(page models.Page) error {
	return r.db.Create(&page).Error
}

func (r *pageRepository) CreatePageWithTx(page models.Page, tx *gorm.DB) error {
	return tx.Create(&page).Error
}

func (r *pageRepository) GetPageByID(id uint) (models.Page, error) {
	var page models.Page
	err := r.db.Where("id = ?", id).First(&page).Error
	return page, err
}

func (r *pageRepository) UpdatePage(page models.Page) error {
	return r.db.Save(&page).Error
}

func (r *pageRepository) UpdatePageWithTx(page models.Page, tx *gorm.DB) error {
	return tx.Save(&page).Error
}

func (r *pageRepository) DeletePage(id uint) error {
	return r.db.Where("id = ?", id).Delete(&models.Page{}).Error
}

func (r *pageRepository) DeletePageWithTx(id uint, tx *gorm.DB) error {
	return tx.Where("id = ?", id).Delete(&models.Page{}).Error
}

func (r *pageRepository) GetAllPages() ([]models.Page, error) {
	var pages []models.Page
	err := r.db.Find(&pages).Error
	return pages, err
}
