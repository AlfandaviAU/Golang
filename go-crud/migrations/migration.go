package main

import (
	"log"

	"github.com/alfandaviau/go-crud/initializers"
	"github.com/alfandaviau/go-crud/models"
)

func init() {
	initializers.Loadenv()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Error cant init model post")
	}
}
