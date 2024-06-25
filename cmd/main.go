package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"os"
	"webcms/cct"
	"webcms/controllers"
	"webcms/models"
	"webcms/repositories"
	"webcms/services"
)

func main() {
	// Инициализация логгера
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	// Копирование кода проекта в файл output.txt
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

	// Маршрут для корневого URL
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Server is running!")
	})

	// Маршруты
	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)

	e.POST("/users", userController.CreateUser)
	e.GET("/users/:id", userController.GetUserByID)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)
	e.GET("/users", userController.GetAllUsers)

	e.POST("/posts", contentController.CreatePost)
	e.GET("/posts/:id", contentController.GetPostByID)
	e.PUT("/posts/:id", contentController.UpdatePost)
	e.DELETE("/posts/:id", contentController.DeletePost)
	e.GET("/posts", contentController.GetAllPosts)

	e.POST("/pages", contentController.CreatePage)
	e.GET("/pages/:id", contentController.GetPageByID)
	e.PUT("/pages/:id", contentController.UpdatePage)
	e.DELETE("/pages/:id", contentController.DeletePage)
	e.GET("/pages", contentController.GetAllPages)

	// Запуск сервера
	sugar.Infof("Starting server on address: %s", ":8080")
	e.Logger.Fatal(e.Start(":8080"))
}
