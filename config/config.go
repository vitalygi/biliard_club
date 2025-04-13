package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db     DbConfig
	JWT    JWTConfig
	Server ServerConfig
}

type ServerConfig struct {
	Port string
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
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
	}
}
