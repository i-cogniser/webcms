package services

import (
	"golang.org/x/crypto/bcrypt"
	"webcms/models"
	"webcms/repositories"
)

type AuthService interface {
	Register(user models.User) error
	Login(email, password string) (models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Register(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepository.CreateUser(user)
}

func (s *authService) Login(email, password string) (models.User, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *authService) FindByEmail(email string) (*models.User, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user.Email == "" {
		return nil, nil
	}
	return &user, nil
}
