// repositories/post_repository.go
package repositories

import (
	"github.com/jinzhu/gorm"
	"webcms/models"
)

type PostRepository interface {
	CreatePost(post models.Post) error
	GetPostByID(id uint) (models.Post, error)
	UpdatePost(post models.Post) error
	DeletePost(id uint) error
	GetAllPosts() ([]models.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) CreatePost(post models.Post) error {
	return r.db.Create(&post).Error
}

func (r *postRepository) GetPostByID(id uint) (models.Post, error) {
	var post models.Post
	err := r.db.Where("id = ?", id).First(&post).Error
	return post, err
}

func (r *postRepository) UpdatePost(post models.Post) error {
	return r.db.Save(&post).Error
}

func (r *postRepository) DeletePost(id uint) error {
	return r.db.Where("id = ?", id).Delete(&models.Post{}).Error
}

func (r *postRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Find(&posts).Error
	return posts, err
}
