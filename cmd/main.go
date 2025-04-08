package main

import (
	"biliard_club/config"
	"biliard_club/internal/auth"
	"biliard_club/internal/user"
	"biliard_club/pkg/db"
	"biliard_club/pkg/middleware"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
)

// TODO
// add https and get port from env
// collect logs in middlewares
// testing

func main() {
	configureLogging()
	conf := config.LoadConfig()
	database := db.NewDb(&conf.Db)
	userRepo := user.NewRepository(database)

	userService := user.NewService(userRepo)
	authService := auth.NewAuthService(userRepo)
	router := gin.Default()

	router.Use(middleware.CORS())

	user.NewHandler(router, user.HandlerDeps{Service: userService})
	auth.NewAuthHandler(router, auth.HandlerDeps{
		JWTConfig: &conf.JWT,
		Service:   authService,
	})

	if err := http.ListenAndServe("localhost:2852", router); err != nil {
		slog.Error("router got fatal error. exiting",
			"error", err.Error())
		panic(err)
	}
}

func configureLogging() {

	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: slog.LevelDebug}))

	slog.SetDefault(logger)
}
