package rest

import (
	"biliard_club/domain"
	"biliard_club/internal/rest/middleware"
	"biliard_club/internal/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandlerDeps struct {
	UserService *user.Service
}
type UserHandler struct {
	UserService *user.Service
}

func NewUserHandler(engine *gin.Engine, deps UserHandlerDeps) {
	handler := UserHandler{
		UserService: deps.UserService,
	}
	protected := engine.Group("/user")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/create", handler.CreateUser)
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var u domain.User
	if err := c.BindJSON(&u); err == nil {
		if _, err = h.UserService.Create(&u); err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}
