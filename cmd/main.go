package main

import (
	"biliard_club/config"
	"biliard_club/internal/repository/postgres"
	"biliard_club/internal/rest"
	"biliard_club/internal/rest/middleware"
	"biliard_club/internal/service/auth"
	"biliard_club/internal/service/table"
	"biliard_club/internal/service/user"
	"biliard_club/pkg/db"
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
	userRepo := postgres.NewUserRepository(database)
	tableRepo := postgres.NewTableRepository(database)

	userService := user.NewService(userRepo)
	authService := auth.NewAuthService(userRepo)
	tableService := table.NewService(tableRepo)

	router := gin.Default()

	router.Use(middleware.CORS())

	rest.NewTableHandler(router, rest.TableHandlerDeps{
		TableService: tableService})
	rest.NewUserHandler(router, rest.UserHandlerDeps{
		UserService: userService})
	rest.NewAuthHandler(router, rest.AuthHandlerDeps{
		JWTConfig: &conf.JWT,
		Service:   authService})

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
