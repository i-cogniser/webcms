// controllers/user_controller.go
package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"webcms/models"
	"webcms/services"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (controller *UserController) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := controller.UserService.CreateUser(*user)
	if err != nil {
		// Возвращаем HTTP-код 500 (внутренняя ошибка сервера) и сообщение об ошибке
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "User created successfully")
}

func (controller *UserController) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := controller.UserService.GetUserByID(uint(id))
	if err != nil {
		// Возвращаем HTTP-код 500 (внутренняя ошибка сервера) и сообщение об ошибке
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user.ID = uint(id)

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = controller.UserService.UpdateUser(*user)
	if err != nil {
		// Возвращаем HTTP-код 500 (внутренняя ошибка сервера) и сообщение об ошибке
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "User updated successfully")
}

func (controller *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = controller.UserService.DeleteUser(uint(id))
	if err != nil {
		// Возвращаем HTTP-код 500 (внутренняя ошибка сервера) и сообщение об ошибке
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "User deleted successfully")
}

func (controller *UserController) GetAllUsers(c echo.Context) error {
	users, err := controller.UserService.GetAllUsers()
	if err != nil {
		// Возвращаем HTTP-код 500 (внутренняя ошибка сервера) и сообщение об ошибке
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}
