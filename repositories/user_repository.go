package repositories

import (
	"github.com/jinzhu/gorm"
	"webcms/models"
)

type UserRepository interface {
	CreateUser(user models.User) error
	CreateUserWithTx(user models.User, tx *gorm.DB) error // Подтверждение наличия метода CreateUserWithTx
	GetUserByID(id uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(user models.User) error
	UpdateUserWithTx(user models.User, tx *gorm.DB) error
	DeleteUser(id uint) error
	DeleteUserWithTx(id uint, tx *gorm.DB) error
	GetAllUsers() ([]models.User, error)
	Count() (int, error)                                  // Добавлен новый метод для подсчета количества пользователей
	FindByUsername(username string) (*models.User, error) // Добавляем метод для поиска пользователя по имени
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) CreateUserWithTx(user models.User, tx *gorm.DB) error {
	return tx.Create(&user).Error
}

func (r *userRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) UpdateUser(user models.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) UpdateUserWithTx(user models.User, tx *gorm.DB) error {
	return tx.Save(&user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Where("id = ?", id).Delete(&models.User{}).Error
}

func (r *userRepository) DeleteUserWithTx(id uint, tx *gorm.DB) error {
	return tx.Where("id = ?", id).Delete(&models.User{}).Error
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

// Count Реализация нового метода для подсчета количества пользователей
func (r *userRepository) Count() (int, error) {
	var count int
	if err := r.db.Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// FindByUsername Реализация нового метода для поиска пользователя по имени
func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
