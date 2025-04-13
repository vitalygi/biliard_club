package rest

import (
	"biliard_club/domain"
	"biliard_club/domain/models"
	"biliard_club/internal/rest/middleware"
	"biliard_club/pkg/validation"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserHandlerDeps struct {
	UserService domain.UserService
}
type UserHandler struct {
	UserService domain.UserService
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
	var u models.User

	if err := c.BindJSON(&u); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			c.JSON(http.StatusBadRequest, gin.H{"error": validation.As(validationErrors).Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "data should be a valid json"})
		}
		return
	}
	if _, err := h.UserService.Create(&u); err != nil {
		var domainErr *domain.Error
		if errors.As(err, &domainErr) {
			c.JSON(domainErr.Code, gin.H{"error": domainErr.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
}
