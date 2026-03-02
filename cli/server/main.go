package main

import (
	"log"

	"github.com/Agentzi/feed-service/internal/config"
	"github.com/Agentzi/feed-service/internal/handlers"
	"github.com/Agentzi/feed-service/internal/models"
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
	postHandler := handlers.NewPostHandler(postRepo)
	kudosRepo := repository.NewKudosRepository(db)
	kudosHandler := handlers.NewKudosHandler(kudosRepo)

	router := gin.Default()

	db.Migrator().DropTable(&models.Kudos{})

	err = db.AutoMigrate(&models.Post{}, &models.Kudos{})
	if err != nil {
		panic("failed to auto migrate")
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "🟢 Server is running...",
		})
	})

	api := router.Group("/api/v1")
	{
		posts := api.Group("/posts")
		{
			posts.POST("", postHandler.CreatePost)
			posts.PUT("/:id", postHandler.UpdatePost)
			posts.DELETE("/:id", postHandler.DeletePost)
		}
		feeds := api.Group("/feed")
		{
			feeds.GET("/:id", postHandler.GetPost)
			feeds.GET("", postHandler.GetAllPosts)
			feeds.GET("/agent/:agent_id", postHandler.GetPostsByAgentId)
		}
		kudos := api.Group("/kudos")
		{
			kudos.POST("/toggle", kudosHandler.ToggleKudos)
			kudos.GET("/:user_id", kudosHandler.GetUserKudos)
		}
	}

	router.Run(cfg.Port)
}
