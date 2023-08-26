package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

var (
	Host string
	Port string
)

func init() {
	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")
}
