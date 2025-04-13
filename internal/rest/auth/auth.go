package auth

import (
	"biliard_club/config"
	"biliard_club/domain"
	"biliard_club/pkg/jwt"
	"biliard_club/pkg/validation"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type AuthHandlerDeps struct {
	*config.JWTConfig
	domain.AuthService
}
type AuthHandler struct {
	*config.JWTConfig
	AuthService domain.AuthService
}

func NewAuthHandler(engine *gin.Engine, deps AuthHandlerDeps) {
	handler := AuthHandler{
		AuthService: deps.AuthService,
		JWTConfig:   deps.JWTConfig,
	}
	engine.POST(
		"/auth/login",
		handler.Login,
	)
	engine.POST(
		"/auth/register",
		handler.Register,
	)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginReq LoginRequest

	if err := c.BindJSON(&loginReq); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			c.JSON(http.StatusBadRequest, gin.H{"error": validation.As(validationErrors).Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "data should be a valid json"})
		}
		return
	}

	user, err := h.AuthService.Login(loginReq.Phone, loginReq.Password)
	if err != nil {
		var domainErr *domain.Error
		if errors.As(err, &domainErr) {
			c.JSON(domainErr.Code, gin.H{"error": domainErr.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	token, err := jwt.NewJWT(h.JWTConfig.Secret).Create(jwt.Data{Phone: user.Phone})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("Authorization", token, 3600*24, "/", "", false, true)

	c.JSON(http.StatusOK, nil)
}

func (h *AuthHandler) Register(c *gin.Context) {

	var registerReq RegisterRequest

	if err := c.BindJSON(&registerReq); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			c.JSON(http.StatusBadRequest, gin.H{"error": validation.As(validationErrors).Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "data should be a valid json"})
		}
		return
	}

	user, err := h.AuthService.Register(registerReq.Phone, registerReq.Password, registerReq.Name)
	if err != nil {
		var domainErr *domain.Error
		if errors.As(err, &domainErr) {
			c.JSON(domainErr.Code, gin.H{"error": domainErr.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	token, err := jwt.NewJWT(h.JWTConfig.Secret).Create(jwt.Data{Phone: user.Phone})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	c.SetCookie("Authorization", token, 3600*24, "/", "", false, true) // 1 день, HttpOnly

	c.JSON(http.StatusOK, nil)
}
