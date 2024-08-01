package routes

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"path/filepath"
	"webcms/controllers"
	"webcms/middlewares"
	"webcms/repositories"
	"webcms/services"
)

func InitRoutes(e *echo.Echo, authController *controllers.AuthController, userController *controllers.UserController, userRepository repositories.UserRepository, authService services.AuthService, tokenRepository repositories.TokenRepository, sugar *zap.SugaredLogger, absFrontendPath string) {
	// Роуты для аутентификации
	e.POST("/api/register", func(c echo.Context) error {
		sugar.Infof("Handling /api/register")
		return authController.Register(c)
	})
	e.POST("/api/login", authController.Login)
	e.POST("/api/refresh-token", authController.RefreshToken)   // Обновление токена
	e.POST("/api/revoke-token/:id", authController.RevokeToken) // Отзыв токена

	// Защищенные маршруты для пользователей
	userGroup := e.Group("/api/users", middlewares.JWTMiddleware(authService))
	userGroup.POST("", userController.CreateUser)
	userGroup.GET("/:id", userController.GetUserByID)
	userGroup.PUT("/:id", userController.UpdateUser)
	userGroup.DELETE("/:id", userController.DeleteUser)
	userGroup.GET("", userController.GetAllUsers)

	// Защищенный маршрут
	protectedGroup := e.Group("/api/protected", middlewares.JWTMiddleware(authService))
	protectedGroup.GET("", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwt.RegisteredClaims)
		return c.JSON(http.StatusOK, claims)
	})

	// Роуты для статистики
	e.GET("/api/users/count", func(c echo.Context) error {
		count, err := userRepository.Count()
		if err != nil {
			sugar.Errorf("Failed to get users count: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get users count"})
		}
		return c.JSON(http.StatusOK, map[string]int{"count": count})
	})

	e.GET("/api/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
	})

	// Роуты для управления токенами
	tokenGroup := e.Group("/api/tokens", middlewares.JWTMiddleware(authService))
	tokenGroup.GET("/refresh", authController.RefreshToken) // Обновление токена
	tokenGroup.DELETE("/:id", authController.RevokeToken)   // Отзыв токена

	// Обработчик для корневого URL
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Статические файлы
	e.Static("/", filepath.Join(absFrontendPath, "public"))

	// Обработчик для всех маршрутов, которые не являются API-запросами
	e.GET("/*", func(c echo.Context) error {
		sugar.Infof("Handling route: %s", c.Request().URL.Path)
		return c.File(filepath.Join(absFrontendPath, "index.html"))
	})
}
