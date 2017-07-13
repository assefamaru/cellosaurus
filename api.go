package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFile("/favicon.ico", "./lib/favicon.ico")

	router.GET("/", func(c *gin.Context) {
		message := "Welcome to Cellosaurus API, see: https://github.com/assefamaru/cellosaurus-api"
		c.String(http.StatusOK, message)
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/", APIVersionEndpoint)
		v1.GET("/cell_lines/", IndexCell)
		v1.GET("/cell_lines/:id", ShowCell)
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
