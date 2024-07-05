package repositories

import (
	"webcms/models"

	"github.com/jinzhu/gorm"
)

type PostRepository interface {
	CreatePost(post models.Post) error
	CreatePostWithTx(post models.Post, tx *gorm.DB) error
	GetPostByID(id uint) (models.Post, error)
	UpdatePost(post models.Post) error
	UpdatePostWithTx(post models.Post, tx *gorm.DB) error
	DeletePost(id uint) error
	DeletePostWithTx(id uint, tx *gorm.DB) error
	GetAllPosts() ([]models.Post, error)
	Count() (int, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) CreatePost(post models.Post) error {
	return r.db.Create(&post).Error
}

func (r *postRepository) CreatePostWithTx(post models.Post, tx *gorm.DB) error {
	return tx.Create(&post).Error
}

func (r *postRepository) GetPostByID(id uint) (models.Post, error) {
	var post models.Post
	err := r.db.Where("id = ?", id).First(&post).Error
	return post, err
}

func (r *postRepository) UpdatePost(post models.Post) error {
	return r.db.Save(&post).Error
}

func (r *postRepository) UpdatePostWithTx(post models.Post, tx *gorm.DB) error {
	return tx.Save(&post).Error
}

func (r *postRepository) DeletePost(id uint) error {
	return r.db.Where("id = ?", id).Delete(&models.Post{}).Error
}

func (r *postRepository) DeletePostWithTx(id uint, tx *gorm.DB) error {
	return tx.Where("id = ?", id).Delete(&models.Post{}).Error
}

func (r *postRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Find(&posts).Error
	return posts, err
}

func (r *postRepository) Count() (int, error) {
	var count int
	err := r.db.Model(&models.Post{}).Count(&count).Error
	return count, err
}
