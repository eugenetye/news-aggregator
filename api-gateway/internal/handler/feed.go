package handler

import (
	"api-gateway/internal/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"log"
	"os"

	redisv9 "github.com/redis/go-redis/v9"
)

func GetTodayFeed(c *gin.Context) {
	today := time.Now().Format("2006-01-02")
	key := "feed:" + today

	log.Println("🔍 Connecting to Redis at:", os.Getenv("REDIS_ADDR"))
	log.Println("📅 Trying to fetch key:", key)

	val, err := redis.Rdb.Get(redis.Ctx, key).Result()
	if err != nil {
		if err == redisv9.Nil {
			log.Println("⚠️ Redis key not found:", key)
		} else {
			log.Println("❌ Redis GET error:", err)
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Feed not found for today"})
		return
	}

	log.Println("✅ Redis GET success. Data length:", len(val))
	c.Data(http.StatusOK, "application/json", []byte(val))
}
