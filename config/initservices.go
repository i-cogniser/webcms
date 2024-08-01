package config

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"webcms/controllers"
	"webcms/repositories"
	"webcms/services"
)

type Services struct {
	UserRepository  repositories.UserRepository
	TokenRepository repositories.TokenRepository
	AuthService     services.AuthService
	UserService     services.UserService
	AuthController  *controllers.AuthController
	UserController  *controllers.UserController
}

func InitServices(db *gorm.DB, sugar *zap.SugaredLogger) *Services {
	userRepository := repositories.NewUserRepository(db)
	tokenRepository := repositories.NewTokenRepository(db)
	authService := services.NewAuthService(userRepository, tokenRepository, db, sugar)
	userService := services.NewUserService(userRepository, db)

	authController := controllers.NewAuthController(authService, sugar)
	userController := controllers.NewUserController(userService, authService, sugar)

	return &Services{
		UserRepository:  userRepository,
		TokenRepository: tokenRepository,
		AuthService:     authService,
		UserService:     userService,
		AuthController:  authController,
		UserController:  userController,
	}
}
