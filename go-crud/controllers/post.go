package controllers

import (
	"log"
	"net/http"

	"github.com/alfandaviau/go-crud/initializers"
	"github.com/alfandaviau/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	res := initializers.DB.Create(&post)

	if res.Error != nil {
		log.Panic(res.Error)
		c.Status(400)
		return
	}

	// Return it
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	var posts models.Post
	initializers.DB.Find(&posts, id)

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"post": posts,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.Find(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.
	Post{Body: body.Body, Title: body.Title})

	// Respond with it
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off url
	id := c.Param("id")

	var post models.Post
	initializers.DB.Find(&post, id)

	initializers.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{
		"deleted": "ok",
	})
}
