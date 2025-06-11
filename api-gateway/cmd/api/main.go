package main

import (
	"api-gateway/internal/handler"
	"api-gateway/internal/redis"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load(".env") // Safe for local only
	if err != nil {
		log.Println("No .env file found (expected for Cloud Run)")
	}

	redis.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	r := gin.Default()

	r.GET("/feed/today", handler.GetTodayFeed)
	r.GET("/alexa/briefing", handler.GetAlexaBriefing)

	log.Printf("🔈 API Gateway listening on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
