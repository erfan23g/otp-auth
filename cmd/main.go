package main

import (
	"otp-auth/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/otp/send", handler.SendOTP)
		v1.POST("/otp/verify", handler.VerifyOTP)
	}

	r.Run(":8080")
}
