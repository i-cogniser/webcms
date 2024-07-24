package controllers

import (
	"go.uber.org/zap"
	"net/http"
	"webcms/models"
	"webcms/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService services.AuthService
	Logger      *zap.SugaredLogger
}

func NewAuthController(authService services.AuthService, Logger *zap.SugaredLogger) *AuthController {
	return &AuthController{AuthService: authService, Logger: Logger}
}

func (controller *AuthController) Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := controller.AuthService.Register(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "User registered successfully")
}

func (controller *AuthController) Login(c echo.Context) error {
	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := controller.AuthService.Login(input.Email, input.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	token, err := controller.AuthService.GenerateJWT(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func (controller *AuthController) RefreshToken(c echo.Context) error {
	currentToken := c.Request().Header.Get("Authorization")
	if currentToken == "" {
		return c.JSON(http.StatusUnauthorized, "No token provided")
	}

	newToken, err := controller.AuthService.RefreshToken(currentToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"token": newToken})
}

func (controller *AuthController) RevokeToken(c echo.Context) error {
	tokenID := c.Param("id")
	err := controller.AuthService.RevokeToken(tokenID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
