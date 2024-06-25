package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath" // добавленный импорт для работы с путями
	"webcms/cct"
	"webcms/controllers"
	"webcms/middlewares" // добавленный импорт для вашего middleware
	"webcms/models"
	"webcms/repositories"
	"webcms/services"
)

func main() {
	// Инициализация логгера
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	// Полное копирование кода проекта в файл output.txt
	cct.CcT()
	// опциональное копирование кода проекта в файл output.txt
	cct.CopyCode()
	// Копирование структуры проекта в файл project_structure.md
	cct.GenTree()
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
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Page{})

	// Инициализация репозиториев
	userRepository := repositories.NewUserRepository(db)
	postRepository := repositories.NewPostRepository(db)
	pageRepository := repositories.NewPageRepository(db)

	// Инициализация сервисов
	authService := services.NewAuthService(userRepository)
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

	// Использование вашего middleware для обработки ошибок
	e.Use(middlewares.ErrorHandler)

	// Маршрут для корневого URL
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running!")
	})

	// Добавление маршрута для статических файлов
	e.Static("/static", "static")

	// Регистрируем обработчик для favicon.ico
	e.File("/favicon.ico", filepath.Join("static", "favicon.ico"))

	// Маршруты
	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)

	// Защищенные маршруты
	userGroup := e.Group("/users", middlewares.AuthMiddleware)
	userGroup.POST("", userController.CreateUser)
	userGroup.GET("/:id", userController.GetUserByID)
	userGroup.PUT("/:id", userController.UpdateUser)
	userGroup.DELETE("/:id", userController.DeleteUser)
	userGroup.GET("", userController.GetAllUsers)

	postGroup := e.Group("/posts", middlewares.AuthMiddleware)
	postGroup.POST("", contentController.CreatePost)
	postGroup.GET("/:id", contentController.GetPostByID)
	postGroup.PUT("/:id", contentController.UpdatePost)
	postGroup.DELETE("/:id", contentController.DeletePost)
	postGroup.GET("", contentController.GetAllPosts)

	pageGroup := e.Group("/pages", middlewares.AuthMiddleware)
	pageGroup.POST("", contentController.CreatePage)
	pageGroup.GET("/:id", contentController.GetPageByID)
	pageGroup.PUT("/:id", contentController.UpdatePage)
	pageGroup.DELETE("/:id", contentController.DeletePage)
	pageGroup.GET("", contentController.GetAllPages)

	// Запуск сервера
	sugar.Infof("Starting server on address: %s", ":8080")
	e.Logger.Fatal(e.Start(":8080"))
}
