package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func LoginRateLimiter(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ip := c.ClientIP()
		key := "login:" + ip
		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			c.JSON(500, gin.H{"error": "redis is unavailable"})
			c.Abort()
			return
		}
		if count == 1 {
			err := rdb.Expire(ctx, key, time.Minute).Err()
			if err != nil {
				c.JSON(500, gin.H{
					"error": "redis unavailable",
				})
				c.Abort()
				return
			}
		}
		if count > 5 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many login attempts"})
			c.Abort()
			return
		}
		c.Next()
	}
}
