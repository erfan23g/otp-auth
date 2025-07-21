package middleware

import (
	"net/http"
	"time"

	"otp-auth/config"

	"github.com/gin-gonic/gin"
)

const (
    Limit     = 10 
    TimeFrame = time.Minute
)

func RateLimiter() gin.HandlerFunc {
    return func(c *gin.Context) {
        clientIP := c.ClientIP()
        key := "rate_limit:" + clientIP

		
        count, err := config.RDB.Incr(c, key).Result()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
            c.Abort()
            return
        }

        
		
        if count == 1 {
            config.RDB.Expire(c, key, TimeFrame)
        }

        
		
        if count > Limit {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests. Please try again later."})
            c.Abort()
            return
        }

        c.Next()
    }
}