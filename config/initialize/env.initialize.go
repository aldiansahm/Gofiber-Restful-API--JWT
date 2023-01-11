package initialize

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	EnvironmentVariables()
}

func EnvironmentVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error load .env file")
	}
}
