package main

import (
	"biliard_club/domain"
	"biliard_club/pkg/db"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	database := db.ConnectDb(os.Getenv("DSN"))

	err = database.AutoMigrate(&domain.User{}, &domain.Table{}, &domain.Game{})
	if err != nil {
		panic(err)
	}
}
