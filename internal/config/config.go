package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL 	string
	RedisURL 		string
	Port 			string
	BaseURL 		string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No se encontró el archivo .env, usando variables del sistema")
	}

	return &Config {
		DatabaseURL: 	os.Getenv("DATABASE_URL"),
		RedisURL: 		os.Getenv("REDIS_URL"),
		Port: 			os.Getenv("PORT"),
		BaseURL: 		os.Getenv("BASE_URL"),
	}
}