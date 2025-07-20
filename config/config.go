package config

import (
	"os"

	"otp-auth/pkg/logger"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

func Init() {
	err := godotenv.Load()
	if err != nil {
		logger.Warn("Could not load .env file")
	}

	// Postgres
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		logger.Fatal("failed to connect database: " + err.Error())
	}
	DB = db
	logger.Info("Connected to PostgreSQL")

	// Redis
	RDB = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
	logger.Info("Redis initialized")
}
