package main

import (
	"log"
	"aggregator/internal/orchestrator"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// Load .env file if it exists (for local development)
    if err := godotenv.Load(".env"); err != nil {
        log.Println("No .env file found (expected in production)")
    }
    
    // Read API keys from environment variables
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        log.Fatal("API_KEY environment variable is required")
    }
    
	outputPath := "daily_feed.json"
	if err := orchestrator.RunAggregation(outputPath); err != nil {
		log.Fatalf("Aggregation failed: %v", err)
	}
}