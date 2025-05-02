package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RateLimiter(limit int, window time.Duration) gin.HandlerFunc {
	requests := make(map[string]int64)
	lastPurge := time.Now()

	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		// Периодическая очистка старых записей
		if now.Sub(lastPurge) > time.Minute {
			requests = make(map[string]int64)
			lastPurge = now
		}

		if count, exists := requests[ip]; exists && count >= int64(limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			return
		}

		requests[ip]++
		c.Next()
	}
}
