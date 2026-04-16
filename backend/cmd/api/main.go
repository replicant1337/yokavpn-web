package main

import (
	"log"
	"os"

	// "yokavpn-web-backend/internal/database"
	"yokavpn-web-backend/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize Database
	// database.InitDB() // Uncomment when DB is running

	r := gin.Default()

	// Routes
	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.POST("/subscriptions", handlers.CreateSubscription)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
