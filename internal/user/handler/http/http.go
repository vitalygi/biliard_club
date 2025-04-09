package http

import (
	"biliard_club/internal/user"
	"biliard_club/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerDeps struct {
	Service *user.Service
}
type Handler struct {
	*user.Service
}

func NewHandler(engine *gin.Engine, deps HandlerDeps) {
	handler := Handler{
		Service: deps.Service,
	}
	protected := engine.Group("/user")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/create", handler.CreateUser)
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user user.User
	if err := c.BindJSON(&user); err == nil {
		if _, err = h.Create(&user); err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}
