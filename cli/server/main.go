package main

import (
	"log"

	"github.com/Agentzi/feed-service/internal/config"
	"github.com/Agentzi/feed-service/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("No .env file found, relying on environment variables: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.DatabaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	postRepo := repository.NewPostRepository(db)
	_ = postRepo

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "🟢 Server is running...",
		})
	})

	router.Run()
}
