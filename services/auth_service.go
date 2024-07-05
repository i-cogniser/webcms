package services

import (
	"errors"
	"os"
	"time"
	"webcms/models"
	"webcms/repositories"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user models.User) error
	RegisterWithTx(user models.User, tx *gorm.DB) error
	Login(email, password string) (*models.User, error)
	GenerateJWT(user *models.User) (string, error)
	GenerateJWTWithTx(user *models.User, tx *gorm.DB) (string, error)
	RefreshToken(tokenString string) (string, error)
	RevokeToken(tokenID string) error
}

type authService struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
	db        *gorm.DB
}

func NewAuthService(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository, db *gorm.DB) AuthService {
	return &authService{userRepo: userRepo, tokenRepo: tokenRepo, db: db}
}

func (s *authService) Register(user models.User) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := s.RegisterWithTx(user, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *authService) RegisterWithTx(user models.User, tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.CreateUserWithTx(user, tx)
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
	return &user, nil
}

func (s *authService) GenerateJWT(user *models.User) (string, error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	token, err := s.GenerateJWTWithTx(user, tx)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()
	return token, nil
}

func (s *authService) GenerateJWTWithTx(user *models.User, tx *gorm.DB) (string, error) {
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

	err = s.tokenRepo.SaveTokenWithTx(models.Token{
		UserID:    user.ID,
		Token:     tokenString,
		ExpiresAt: expirationTime,
	}, tx)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *authService) RefreshToken(tokenString string) (string, error) {
	token, err := s.tokenRepo.GetToken(tokenString)
	if err != nil || token.ExpiresAt.Before(time.Now()) {
		return "", errors.New("invalid or expired token")
	}

	user, err := s.userRepo.GetUserByID(token.UserID)
	if err != nil {
		return "", err
	}

	return s.GenerateJWT(&user)
}

func (s *authService) RevokeToken(tokenID string) error {
	return s.tokenRepo.DeleteToken(tokenID)
}
