package repository

import (
	"fmt"
	"log"
	"os"
	"otp-auth/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := os.Getenv("DB_DSN")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate DB: ", err)
	}

	fmt.Println("Database connection successful.")
}
func FirstOrCreate(user *model.User, newUser model.User) *gorm.DB {
	result := DB.FirstOrCreate(user, newUser)
	if result.Error != nil {
		log.Printf("Error finding or creating user: %v", result.Error)
	}
	return result
}
