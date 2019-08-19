package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var PORT string

func init()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("HTTP_PORT")
	log.Print(PORT)
}