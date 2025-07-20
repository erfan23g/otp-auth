package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"otp-auth/config"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func SetOTP(phone string) (string, error) {
	code := generateOTP()
	err := config.RDB.Set(ctx, phone, code, 2*time.Minute).Err()
	return code, err
}

func VerifyOTP(phone, code string) bool {
	val, err := config.RDB.Get(ctx, phone).Result()
	if err == redis.Nil || err != nil {
		return false
	}
	return val == code
}
