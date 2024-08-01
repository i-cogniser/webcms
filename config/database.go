package config

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
	"webcms/models"

	_ "github.com/lib/pq"
)

func GetDBParams() (string, string, string, string, string) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	return dbHost, dbPort, dbUser, dbPassword, dbName
}

func GetDBURL(sugar *zap.SugaredLogger) string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	sugar.Infof("Connecting to database: %s", dbURL)
	return dbURL
}

func WaitForDB(dbURL string, sugar *zap.SugaredLogger) error {
	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err == nil {
			sqlDB, _ := db.DB()
			defer sqlDB.Close()
			if err := sqlDB.Ping(); err == nil {
				return nil
			}
		}
		sugar.Infof("Database not ready, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
	return fmt.Errorf("database not ready after 10 attempts: %w", err)
}

func CreateDatabaseAndUser(sugar *zap.SugaredLogger) error {
	dbHost, dbPort, dbUser, dbPassword, dbName := GetDBParams()
	dsn := fmt.Sprintf("host=%s port=%s user=postgres password=%s sslmode=disable", dbHost, dbPort, dbPassword)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer db.Close()

	// Проверка существования базы данных
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname='%s')", dbName)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	if !exists {
		// Создание базы данных
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
		sugar.Infof("Database %s created successfully", dbName)
	} else {
		sugar.Infof("Database %s already exists", dbName)
	}

	// Проверка существования пользователя
	var userExists bool
	userQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_roles WHERE rolname='%s')", dbUser)
	err = db.QueryRow(userQuery).Scan(&userExists)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}

	if !userExists {
		// Создание пользователя
		_, err = db.Exec(fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s'", dbUser, dbPassword))
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		sugar.Infof("User %s created successfully", dbUser)
	}

	// Назначение прав пользователю
	_, err = db.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s", dbName, dbUser))
	if err != nil {
		return fmt.Errorf("failed to grant privileges to user: %w", err)
	}

	return nil
}

func InitDB(dbURL string, sugar *zap.SugaredLogger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		sugar.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func MigrateDB(db *gorm.DB, sugar *zap.SugaredLogger) error {
	sugar.Info("Starting database migration...")
	if err := db.AutoMigrate(&models.User{}, &models.Token{}); err != nil {
		sugar.Errorf("Failed to migrate database: %v", err)
		return err
	}
	sugar.Info("Database migration completed successfully")
	return nil
}
