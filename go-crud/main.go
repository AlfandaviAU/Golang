package main

import (
	"net/http"

	"github.com/alfandaviau/go-crud/controllers"
	"github.com/alfandaviau/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Loadenv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
