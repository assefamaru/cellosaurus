package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		message := "Welcome to Cellosaurus API, see: https://github.com/assefamaru/cellosaurus-api"
		c.String(http.StatusOK, message)
	})

	v1 := router.Group("/v1")
	{
		router.GET("/", func(c *gin.Context) {
			message := "Welcome to Cellosaurus API, see: https://github.com/assefamaru/cellosaurus-api"
			c.String(http.StatusOK, message)
		})

		v1.GET("/cell_lines", func(c *gin.Context) {})
		v1.GET("/cell_lines/:id", func(c *gin.Context) {})
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    http.StatusBadRequest,
				"message": "Bad Request",
			},
		})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
