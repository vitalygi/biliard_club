package main

import (
	"biliard_club/internal/game"
	"biliard_club/internal/table"
	"biliard_club/internal/user"
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

	err = database.AutoMigrate(&user.User{}, &table.Table{}, &game.Game{})
	if err != nil {
		panic(err)
	}
}
