package services

import (
	"github.com/jinzhu/gorm"
	"webcms/models"
	"webcms/repositories"
)

type UserService interface {
	GetUserByID(id uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id uint) error
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.User) error
}

type userService struct {
	userRepository repositories.UserRepository
	db             *gorm.DB
}

func NewUserService(userRepo repositories.UserRepository, db *gorm.DB) UserService {
	return &userService{userRepo, db}
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *userService) GetUserByEmail(email string) (models.User, error) {
	return s.userRepository.GetUserByEmail(email)
}

func (s *userService) UpdateUser(user models.User) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.userRepository.UpdateUserWithTx(user, tx)
	})
}

func (s *userService) DeleteUser(id uint) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.userRepository.DeleteUserWithTx(id, tx)
	})
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAllUsers()
}

func (s *userService) CreateUser(user models.User) error {
	return execWithTx(s.db, func(tx *gorm.DB) error {
		return s.userRepository.CreateUserWithTx(user, tx)
	})
}
