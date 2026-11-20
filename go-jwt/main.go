package main

import (
	"net/http"

	"github.com/alfandaviau/go-jwt/controllers"
	"github.com/alfandaviau/go-jwt/initializers"
	"github.com/alfandaviau/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Loadenv()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "ok",
		})
	})

	r.GET("/users", controllers.GetAllUser)
	r.GET("/users/:id", controllers.GetSingleUser)
	r.POST("/signup", controllers.SignUp)

	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
