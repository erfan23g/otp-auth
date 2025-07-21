package main

import (
	"otp-auth/config"
	"otp-auth/internal/model"
	"otp-auth/internal/repository"
	"otp-auth/internal/handler"
	"otp-auth/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	config.DB.AutoMigrate(&model.User{})

	repository.Init()

	r := gin.Default()
	r.Use(middleware.RateLimiter())

	v1 := r.Group("/api/v1")
	{
		v1.POST("/otp/send", handler.SendOTP)
		v1.POST("/otp/verify", handler.VerifyOTP)
	}

	r.Run(":8080")
}
