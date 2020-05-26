package main

import (
	"net/http"
	"os"
	"strings"

	"config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitializeEnv()

	PORT := strings.TrimSpace(os.Getenv("PORT"))
	DB := strings.TrimSpace(os.Getenv("DB"))

	config.DBConnection(DB)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
