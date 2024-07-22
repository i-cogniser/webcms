package controllers

import (
	"net/http"
	"webcms/models"
	"webcms/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AuthController struct {
	AuthService services.AuthService
	Logger      *zap.SugaredLogger
}

func NewAuthController(authService services.AuthService, logger *zap.SugaredLogger) *AuthController {
	return &AuthController{AuthService: authService, Logger: logger}
}

func (controller *AuthController) Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		controller.Logger.Errorf("Error binding user: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		controller.Logger.Errorf("Validation error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := controller.AuthService.Register(*user)
	if err != nil {
		controller.Logger.Errorf("Error registering user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	controller.Logger.Infof("User registered successfully: %v", user.Email)
	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully"})
}

func (controller *AuthController) Login(c echo.Context) error {
	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		controller.Logger.Errorf("Error binding login input: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		controller.Logger.Errorf("Validation error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := controller.AuthService.Login(input.Email, input.Password)
	if err != nil {
		controller.Logger.Errorf("Login failed: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	token, err := controller.AuthService.GenerateJWT(user)
	if err != nil {
		controller.Logger.Errorf("Error generating token: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	controller.Logger.Infof("User logged in successfully: %v", user.Email)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func (controller *AuthController) RefreshToken(c echo.Context) error {
	currentToken := c.Request().Header.Get("Authorization")
	if currentToken == "" {
		controller.Logger.Error("No token provided")
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "No token provided"})
	}

	newToken, err := controller.AuthService.RefreshToken(currentToken)
	if err != nil {
		controller.Logger.Errorf("Error refreshing token: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	controller.Logger.Infof("Token refreshed successfully")
	return c.JSON(http.StatusOK, map[string]string{"token": newToken})
}

func (controller *AuthController) RevokeToken(c echo.Context) error {
	tokenID := c.Param("id")
	err := controller.AuthService.RevokeToken(tokenID)
	if err != nil {
		controller.Logger.Errorf("Error revoking token: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	controller.Logger.Infof("Token revoked successfully: %v", tokenID)
	return c.NoContent(http.StatusOK)
}
