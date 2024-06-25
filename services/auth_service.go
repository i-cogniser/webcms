package services

import (
	"errors"
	"os"
	"time"
	"webcms/models"
	"webcms/repositories"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user models.User) error
	Login(email, password string) (*models.User, error)
	GenerateJWT(user *models.User) (string, error)
}

type authService struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
}

func NewAuthService(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository) AuthService {
	return &authService{userRepo, tokenRepo}
}

func (s *authService) Register(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepo.CreateUser(user)
}

func (s *authService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil // Возвращаем указатель на user
}

func (s *authService) GenerateJWT(user *models.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT secret not configured")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   string(user.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	// Save the token in the database
	err = s.tokenRepo.SaveToken(models.Token{
		UserID:    user.ID,
		Token:     tokenString,
		ExpiresAt: expirationTime,
	})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
