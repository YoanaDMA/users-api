package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const message = "Users API starting up... 🚀"

func main() {
	response := fmt.Sprintf("Hello, World!!!, %s", message)
	fmt.Println(response)

	port := os.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "up",
			"service": "users-api",
		})
	})

	fmt.Printf("🚀 Server running on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
