package main

import (
	"album/backend/internal/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("❌ error loading .env file")
	}
	db := config.InitDB()
	defer db.Close()
	r.Run()
}
