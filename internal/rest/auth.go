package rest

import (
	"biliard_club/config"
	"biliard_club/internal/service/auth"
	"biliard_club/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
type RegisterResponse struct {
	Token string `json:"token"`
}

type AuthHandlerDeps struct {
	*config.JWTConfig
	*auth.Service
}
type AuthHandler struct {
	*config.JWTConfig
	AuthService *auth.Service
}

func NewAuthHandler(engine *gin.Engine, deps AuthHandlerDeps) {
	handler := AuthHandler{
		AuthService: deps.Service,
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
		return
	}
	user, err := h.AuthService.Login(loginReq.Phone, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.AuthService.Register(registerReq.Phone, registerReq.Password, registerReq.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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
