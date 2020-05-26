package main

import (
	"net/http"
	"os"

	"config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitializeEnv()
	PORT := os.Getenv("PORT")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
