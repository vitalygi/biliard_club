package main

import (
	"biliard_club/config"
	"biliard_club/internal/auth"
	http3 "biliard_club/internal/auth/handler/http"
	"biliard_club/internal/table"
	httphandler "biliard_club/internal/table/handler/http"
	"biliard_club/internal/user"
	http2 "biliard_club/internal/user/handler/http"
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
	tableRepo := table.NewRepository(database)

	userService := user.NewService(userRepo)
	authService := auth.NewAuthService(userRepo)
	tableService := table.NewService(tableRepo)

	router := gin.Default()

	router.Use(middleware.CORS())

	httphandler.NewHandler(router, httphandler.HandlerDeps{Service: tableService})
	http2.NewHandler(router, http2.HandlerDeps{Service: userService})
	http3.NewAuthHandler(router, http3.HandlerDeps{
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
