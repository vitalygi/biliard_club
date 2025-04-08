package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db  DbConfig
	JWT JWTConfig
}

type JWTConfig struct {
	Secret string
}
type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error while loading config")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		JWT: JWTConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}
