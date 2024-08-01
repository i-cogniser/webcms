package main

import (
	"os"
	"path/filepath"
	"webcms/config"
	"webcms/rendering"
	"webcms/routes"
)

func main() {
	// Инициализация логгера
	sugar := config.InitLogger()
	defer sugar.Sync()

	// Загрузка переменных окружения
	config.LoadEnv(sugar)

	// Создание базы данных и пользователя, если они не существуют
	config.CreateDatabaseAndUser(sugar)

	// Получение URL базы данных и логирование
	dbURL := config.GetDBURL(sugar)

	// Ожидание готовности базы данных
	config.WaitForDB(dbURL, sugar)

	// Инициализация базы данных
	db := config.InitDB(dbURL, sugar)
	defer config.CloseDB(db)

	// Автоматическая миграция моделей
	config.MigrateDB(db, sugar)

	// Инициализация репозиториев, сервисов и контроллеров
	services := config.InitServices(db, sugar)

	// Инициализация Echo
	e := config.SetupEcho(sugar)

	// Настройка рендеринга шаблонов
	rendering.SetupRenderer(e)

	// Статические файлы
	frontendPath := os.Getenv("FRONTEND_PATH")
	if frontendPath == "" {
		frontendPath = "webcms-vue/dist"
	}
	absFrontendPath, err := filepath.Abs(frontendPath)
	if err != nil {
		sugar.Fatalf("Failed to get absolute path for frontend: %v", err)
	}

	// Инициализация маршрутов
	routes.InitRoutes(e, services.AuthController, services.UserController, services.UserRepository, services.AuthService, services.TokenRepository, sugar, absFrontendPath)

	// Периодическая проверка и удаление устаревших токенов
	go config.StartTokenCleanupProcess(services.TokenRepository, sugar)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
