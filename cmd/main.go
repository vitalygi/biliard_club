package main

import (
	"biliard_club/config"
	"biliard_club/internal/repository/postgres"
	"biliard_club/internal/rest"
	restAuth "biliard_club/internal/rest/auth"
	"biliard_club/internal/rest/middleware"
	"biliard_club/internal/service/auth"
	"biliard_club/internal/service/table"
	"biliard_club/internal/service/user"
	"biliard_club/pkg/db"
	"biliard_club/pkg/validation"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	sloggin "github.com/samber/slog-gin"
	"log/slog"
	"net/http"
	"os"
)

// TODO
// logs observer
// testing

func main() {
	conf := config.LoadConfig()

	setGinMode()
	router := gin.New()
	configureLogging(router)
	configureGin(router)

	database := db.NewDb(&conf.Db)
	userRepo := postgres.NewUserRepository(database)
	tableRepo := postgres.NewTableRepository(database)

	userService := user.NewService(userRepo)
	authService := auth.NewAuthService(userRepo)
	tableService := table.NewService(tableRepo)

	router.Use(middleware.CORS())

	rest.NewTableHandler(router, rest.TableHandlerDeps{
		TableService: tableService})
	rest.NewUserHandler(router, rest.UserHandlerDeps{
		UserService: userService})
	restAuth.NewAuthHandler(router, restAuth.AuthHandlerDeps{
		JWTConfig:   &conf.JWT,
		AuthService: authService})

	addr := fmt.Sprintf("%s:%s", "localhost", conf.Server.Port)
	if err := http.ListenAndServe(addr, router); err != nil {
		slog.Error("router got fatal error. exiting",
			"error", err.Error())
		panic(err)
	}
}

func configureLogging(router *gin.Engine) {
	gin.Logger()
	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: slog.LevelDebug}))
	router.Use(sloggin.New(logger))
	slog.SetDefault(logger)
}

func setGinMode() {
	// load gin mode from env
	GIN_MODE := os.Getenv("GIN_MODE")
	if GIN_MODE != "" {
		gin.SetMode(GIN_MODE)
	}
}
func configureGin(router *gin.Engine) {
	// set recovery middleware
	router.Use(gin.Recovery())
	// set gin default validator
	binding.Validator = validation.GinValidator
}
