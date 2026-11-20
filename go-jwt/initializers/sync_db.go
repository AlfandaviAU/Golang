package initializers

import "github.com/alfandaviau/go-jwt/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
