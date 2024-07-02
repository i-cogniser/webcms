package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"webcms/models"
	"webcms/services"
)

type ContentController struct {
	ContentService services.ContentService
}

func NewContentController(contentService services.ContentService) *ContentController {
	return &ContentController{contentService}
}

// Post Handlers

func (controller *ContentController) CreatePost(c echo.Context) error {
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := controller.ContentService.CreatePost(*post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "Post created successfully")
}

func (controller *ContentController) GetPostByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	post, err := controller.ContentService.GetPostByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, post)
}

func (controller *ContentController) UpdatePost(c echo.Context) error {
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := controller.ContentService.UpdatePost(*post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "Post updated successfully")
}

func (controller *ContentController) DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = controller.ContentService.DeletePost(uint(id))
	if err != nil {
		// Возвращаем HTTP-код 500 (внутренняя ошибка сервера) и сообщение об ошибке
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "Post deleted successfully")
}

func (controller *ContentController) GetAllPosts(c echo.Context) error {
	posts, err := controller.ContentService.GetAllPosts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, posts)
}

// Page Handlers

func (controller *ContentController) CreatePage(c echo.Context) error {
	page := new(models.Page)
	if err := c.Bind(page); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := controller.ContentService.CreatePage(*page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "Page created successfully")
}

func (controller *ContentController) GetPageByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	page, err := controller.ContentService.GetPageByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, page)
}

func (controller *ContentController) UpdatePage(c echo.Context) error {
	page := new(models.Page)
	if err := c.Bind(page); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := controller.ContentService.UpdatePage(*page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "Page updated successfully")
}

func (controller *ContentController) DeletePage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = controller.ContentService.DeletePage(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "Page deleted successfully")
}

func (controller *ContentController) GetAllPages(c echo.Context) error {
	pages, err := controller.ContentService.GetAllPages()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, pages)
}
