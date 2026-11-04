package main

import (
	"context"
	"log"
	"net/http"

	"website/internal/config"
	"website/internal/database"
	"website/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()

	database.RunMigrations(cfg.DATABASE_URL)

	dbConn, err := database.Connect(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("Falha ao conectar no banco: %v", err)
	}
	defer dbConn.Close(ctx)

	repository.NewPostRepository(dbConn)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3000")
}
