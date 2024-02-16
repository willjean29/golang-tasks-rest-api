package plugins

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvLoader interface {
	Load(string) error
}
type Env struct {
	TokenSecretKey string
	DBPort         string
	DBHost         string
	DBName         string
	DBUser         string
	DBPassword     string
}

func NewEnv() *Env {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Env{
		TokenSecretKey: os.Getenv("JWT_SECRET_KEY"),
		DBPort:         os.Getenv("DB_PORT"),
		DBHost:         os.Getenv("DB_HOST"),
		DBName:         os.Getenv("DB_NAME"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
	}
}

var Envs = NewEnv()
