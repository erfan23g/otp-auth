package handler

import (
	"fmt"
	"net/http"
	"otp-auth/config"
	"otp-auth/internal/model"
	"otp-auth/internal/repository"
	"otp-auth/internal/service"
	"otp-auth/pkg/logger"

	"github.com/gin-gonic/gin"
)

type SendOTPRequest struct {
	Phone string `json:"phone" binding:"required"`
}

type VerifyOTPRequest struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

func SendOTP(c *gin.Context) {
	fmt.Println("HII")
	var req SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone"})
		return
	}


	var user model.User
	result := repository.FirstOrCreate(&user, model.User{PhoneNumber: req.Phone})
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Could not save the user"})
		return
	}

	code, err := service.SetOTP(req.Phone)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not set OTP"})
		return
	}

	logger.Info("OTP code: " + code)
}

func VerifyOTP(c *gin.Context) {
	var req VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if !service.VerifyOTP(req.Phone, req.Code) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "OTP invalid or expired"})
		return
	}

	
	var user model.User
	result := config.DB.Where("phone_number = ?", req.Phone).First(&user)
	if result.Error != nil {
		user = model.User{PhoneNumber: req.Phone}
		config.DB.Create(&user)
	}

	token, err := service.GenerateJWT(req.Phone)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are authenticated!"})
}
