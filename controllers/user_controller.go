package controllers

import (
	"net/http"
	"strconv"
	"webcms/models"
	"webcms/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserController struct {
	UserService services.UserService
	AuthService services.AuthService
	Logger      *zap.SugaredLogger
}

func NewUserController(userService services.UserService, authService services.AuthService, logger *zap.SugaredLogger) UserController {
	return UserController{UserService: userService, AuthService: authService, Logger: logger}
}

func (controller *UserController) Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		controller.Logger.Errorf("Error binding user: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request"})
	}

	controller.Logger.Infof("User data after binding: %+v", user)

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		controller.Logger.Errorf("Validation error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Validation error"})
	}

	controller.Logger.Infof("User data after validation: %+v", user)

	err := controller.AuthService.Register(*user)
	if err != nil {
		controller.Logger.Errorf("Error registering user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	controller.Logger.Info("User successfully registered")

	return c.JSON(http.StatusOK, "User successfully registered")
}

func (controller UserController) Login(c echo.Context) error {
	loginRequest := new(struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	})

	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := controller.AuthService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	token, err := controller.AuthService.GenerateJWT(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (controller UserController) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := controller.UserService.CreateUser(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "Пользователь успешно создан")
}

func (controller UserController) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := controller.UserService.GetUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (controller UserController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	user.ID = uint(id)

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = controller.UserService.UpdateUser(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "Пользователь успешно обновлен")
}

func (controller UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = controller.UserService.DeleteUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "Пользователь успешно удален")
}

func (controller UserController) GetAllUsers(c echo.Context) error {
	controller.Logger.Info("GetAllUsers method called")
	users, err := controller.UserService.GetAllUsers()
	if err != nil {
		controller.Logger.Errorf("Error fetching users: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	controller.Logger.Infof("Users fetched successfully: %d users", len(users))
	return c.JSON(http.StatusOK, users)
}
