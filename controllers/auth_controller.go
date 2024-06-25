// controllers/auth_controller.go
package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"webcms/models"
	"webcms/services"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
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
		// Возвращаем HTTP-код 500 (внутренняя ошибка сервера) и сообщение об ошибке
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "User registered successfully")
}

func (controller *AuthController) Login(c echo.Context) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := controller.AuthService.Login(input.Email, input.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, user)
}
