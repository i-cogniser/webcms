package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"webcms/cct"
	"webcms/controllers"
	"webcms/middlewares"
	"webcms/models"
	"webcms/rendering"
	"webcms/repositories"
	"webcms/services"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	// Полное копирование кода проекта в файл output.txt
	cct.CcT()
	// опциональное копирование кода проекта в файл output.txt
	cct.CopyCode()
	// Копирование структуры проекта в файл project_structure.md
	cct.GenTree()
	// Инициализация логгера
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	// Загрузка переменных окружения из .env файла
	if err := godotenv.Load(); err != nil {
		sugar.Fatalf("Error loading .env file: %v", err)
	}

	// Формирование DATABASE_URL
	dbURL := os.Getenv("DATABASE_URL")

	// Если DATABASE_URL не установлен в .env, собираем его из отдельных переменных
	if dbURL == "" {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		dbURL = "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName
	}

	// Логирование URL базы данных
	sugar.Infof("Connecting to database: %s", dbURL)

	// Инициализация базы данных
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		sugar.Fatalf("Failed to connect database: %v", err)
		return
	}
	defer db.Close()

	// Автоматическая миграция моделей
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Page{}, &models.Token{})

	// Инициализация репозиториев
	userRepository := repositories.NewUserRepository(db)
	postRepository := repositories.NewPostRepository(db)
	pageRepository := repositories.NewPageRepository(db)
	tokenRepository := repositories.NewTokenRepository(db)

	// Инициализация сервисов
	authService := services.NewAuthService(userRepository, tokenRepository)
	userService := services.NewUserService(userRepository)
	contentService := services.NewContentService(postRepository, pageRepository)

	// Инициализация контроллеров
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	contentController := controllers.NewContentController(contentService)

	// Инициализация Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Использование CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:8081"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Настройка структуры для рендеринга HTML шаблонов
	renderer := &rendering.TemplateRenderer{
		Templates: template.Must(template.ParseGlob(filepath.Join("templates", "*.html"))),
	}
	e.Renderer = renderer

	// Применение middleware для проверки токена на защищенных маршрутах
	// Routes
	e.POST("/api/register", authController.Register)
	e.POST("/api/login", authController.Login)

	// Защищенные маршруты
	userGroup := e.Group("/users", middlewares.JWTMiddleware(authService))
	userGroup.POST("", userController.CreateUser)
	userGroup.GET("/:id", userController.GetUserByID)
	userGroup.PUT("/:id", userController.UpdateUser)
	userGroup.DELETE("/:id", userController.DeleteUser)
	userGroup.GET("", userController.GetAllUsers)

	postGroup := e.Group("/posts", middlewares.JWTMiddleware(authService))
	postGroup.POST("", contentController.CreatePost)
	postGroup.GET("/:id", contentController.GetPostByID)
	postGroup.PUT("/:id", contentController.UpdatePost)
	postGroup.DELETE("/:id", contentController.DeletePost)
	postGroup.GET("", contentController.GetAllPosts)

	pageGroup := e.Group("/pages", middlewares.JWTMiddleware(authService))
	pageGroup.POST("", contentController.CreatePage)
	pageGroup.GET("/:id", contentController.GetPageByID)
	pageGroup.PUT("/:id", contentController.UpdatePage)
	pageGroup.DELETE("/:id", contentController.DeletePage)
	pageGroup.GET("", contentController.GetAllPages)

	// Обработчик для корневого URL
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Статические файлы
	frontendPath := os.Getenv("FRONTEND_PATH")
	if frontendPath == "" {
		frontendPath = "web"
	}

	absFrontendPath, err := filepath.Abs(frontendPath)
	if err != nil {
		sugar.Fatalf("Failed to get absolute path for frontend: %v", err)
	}
	e.Static("/", absFrontendPath)

	// Периодическая проверка и удаление устаревших токенов
	go func() {
		for {
			err := tokenRepository.DeleteExpiredTokens()
			if err != nil {
				sugar.Errorf("Failed to delete expired tokens: %v", err)
			}
			time.Sleep(1 * time.Hour)
		}
	}()

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
