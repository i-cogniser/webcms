package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"webcms/controllers"
	"webcms/middlewares"
	"webcms/models"
	"webcms/rendering"
	"webcms/repositories"
	"webcms/services"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Используем текущую версию GORM
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	// Инициализация логгера
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	fmt.Println("Checkpoint 1: Logger initialized")

	// Загрузка переменных окружения из .env файла
	if err := godotenv.Load(); err != nil {
		sugar.Fatalf("Error loading .env file: %v", err)
	} else {
		sugar.Infof(".env file loaded successfully")
	}
	fmt.Println("Checkpoint 2: Env loaded")

	// Формирование DATABASE_URL
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	}
	fmt.Println("Checkpoint 3: DB URL formed")

	// Логирование URL базы данных
	sugar.Infof("Connecting to database: %s", dbURL)

	// Инициализация базы данных
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		sugar.Fatalf("Failed to connect database: %v", err)
		return
	} else {
		sugar.Infof("Database connected successfully")
	}
	defer db.Close()

	fmt.Println("Checkpoint 4: DB connected")

	// Автоматическая миграция моделей
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Page{}, &models.Token{}).Error; err != nil {
		sugar.Fatalf("Failed to migrate database: %v", err)
	} else {
		sugar.Infof("Database migrated successfully")
	}

	// Инициализация репозиториев
	userRepository := repositories.NewUserRepository(db)
	postRepository := repositories.NewPostRepository(db)
	pageRepository := repositories.NewPageRepository(db)
	tokenRepository := repositories.NewTokenRepository(db)

	// Инициализация сервисов
	authService := services.NewAuthService(userRepository, tokenRepository, db)
	userService := services.NewUserService(userRepository, db)
	contentService := services.NewContentService(postRepository, pageRepository, db)

	// Инициализация контроллеров
	authController := controllers.NewAuthController(authService, sugar)
	userController := controllers.NewUserController(userService, authService, sugar)
	contentController := controllers.NewContentController(contentService)

	// Инициализация Echo
	e := echo.New()

	// Middleware для логирования запросов
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
			method := c.Request().Method
			sugar.Infof("Request received: %s %s", method, path)
			return next(c)
		}
	})

	// Middleware для обработки логов и восстановления
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Использование CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "http://localhost:80", "http://localhost:8080", "http://localhost:8081"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Настройка структуры для рендеринга HTML шаблонов
	renderer := &rendering.TemplateRenderer{
		Templates: template.Must(template.ParseGlob(filepath.Join("templates", "*.html"))),
	}
	e.Renderer = renderer

	// Настройка обработчика ошибок
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == http.StatusNotFound {
				c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found"})
				return
			}
		}
		e.DefaultHTTPErrorHandler(err, c)
	}

	// Роуты для аутентификации
	e.POST("/api/register", authController.Register)
	e.POST("/api/login", authController.Login)
	e.POST("/api/refresh-token", authController.RefreshToken)   // Обновление токена
	e.POST("/api/revoke-token/:id", authController.RevokeToken) // Отзыв токена

	// Защищенные маршруты для пользователей
	userGroup := e.Group("/users", middlewares.JWTMiddleware(authService))
	userGroup.POST("", func(c echo.Context) error { // <-- Здесь добавляем новый маршрут
		var user models.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		existingUser, err := userRepository.FindByUsername(user.Username)
		if err == nil && existingUser != nil {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Username already exists"})
		}

		if err := userRepository.CreateUser(user); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
		}

		return c.JSON(http.StatusCreated, user)
	})
	userGroup.POST("", userController.CreateUser)
	userGroup.GET("/:id", userController.GetUserByID)
	userGroup.PUT("/:id", userController.UpdateUser)
	userGroup.DELETE("/:id", userController.DeleteUser)
	userGroup.GET("", userController.GetAllUsers)

	// Защищенные маршруты для постов
	postGroup := e.Group("/posts", middlewares.JWTMiddleware(authService))
	postGroup.POST("", contentController.CreatePost)
	postGroup.GET("/:id", contentController.GetPostByID)
	postGroup.PUT("/:id", contentController.UpdatePost)
	postGroup.DELETE("/:id", contentController.DeletePost)
	postGroup.GET("", contentController.GetAllPosts)

	// Защищенные маршруты для страниц
	pageGroup := e.Group("/pages", middlewares.JWTMiddleware(authService))
	pageGroup.POST("", contentController.CreatePage)
	pageGroup.GET("/:id", contentController.GetPageByID)
	pageGroup.PUT("/:id", contentController.UpdatePage)
	pageGroup.DELETE("/:id", contentController.DeletePage)
	pageGroup.GET("", contentController.GetAllPages)

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

	e.GET("/api/pages/count", func(c echo.Context) error {
		count, err := pageRepository.Count()
		if err != nil {
			sugar.Errorf("Failed to get pages count: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get users count"})
		}
		return c.JSON(http.StatusOK, map[string]int{"count": count})
	})

	e.GET("/api/posts/count", func(c echo.Context) error {
		count, err := postRepository.Count()
		if err != nil {
			sugar.Errorf("Failed to get posts count: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get users count"})
		}
		return c.JSON(http.StatusOK, map[string]int{"count": count})
	})

	e.GET("/api/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
	})

	// Роуты для управления токенами
	tokenGroup := e.Group("/tokens", middlewares.JWTMiddleware(authService))
	tokenGroup.GET("/refresh", authController.RefreshToken) // Обновление токена
	tokenGroup.DELETE("/:id", authController.RevokeToken)   // Отзыв токена

	// Обработчик для корневого URL
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Статические файлы
	frontendPath := os.Getenv("FRONTEND_PATH")
	if frontendPath == "" {
		frontendPath = "webcms-vue/dist"
	}

	absFrontendPath, err := filepath.Abs(frontendPath)
	if err != nil {
		sugar.Fatalf("Failed to get absolute path for frontend: %v", err)
	}
	e.Static("/", filepath.Join(absFrontendPath, "public"))

	// Обработчик для всех маршрутов, которые не являются API-запросами
	e.GET("/*", func(c echo.Context) error {
		sugar.Infof("Handling route: %s", c.Request().URL.Path)
		return c.File(filepath.Join(absFrontendPath, "index.html"))
	})

	// Периодическая проверка и удаление устаревших токенов
	go func() {
		for {
			sugar.Infof("Starting token cleanup process")
			err := tokenRepository.DeleteExpiredTokens()
			if err != nil {
				sugar.Errorf("Failed to delete expired tokens: %v", err)
			} else {
				sugar.Infof("Expired tokens deleted successfully")
			}
			time.Sleep(1 * time.Hour)
		}
	}()

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
