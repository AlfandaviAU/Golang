package initializers

import (
	"github.com/joho/godotenv"
)

func Loadenv() {
	godotenv.Load()
}
