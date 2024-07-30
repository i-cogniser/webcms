package repositories

import (
	"errors"
	"webcms/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) error
	CreateUserWithTx(user models.User, tx *gorm.DB) error
	GetUserByID(id uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(user models.User) error
	UpdateUserWithTx(user models.User, tx *gorm.DB) error
	DeleteUser(id uint) error
	DeleteUserWithTx(id uint, tx *gorm.DB) error
	GetAllUsers() ([]models.User, error)
	Count() (int, error)
	FindByUsername(username string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user models.User) error {
	var existingUser models.User
	if err := r.db.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
		return errors.New("user with the same username or email already exists")
	} else if !gorm.IsRecordNotFoundError(err) {
		return err
	}
	return r.db.Create(&user).Error
}

func (r *userRepository) CreateUserWithTx(user models.User, tx *gorm.DB) error {
	var existingUser models.User
	if err := tx.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
		return errors.New("user with the same username or email already exists")
	} else if !gorm.IsRecordNotFoundError(err) {
		return err
	}
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

func (r *userRepository) Count() (int, error) {
	var count int
	if err := r.db.Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

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
