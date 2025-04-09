package http

import (
	"biliard_club/config"
	"biliard_club/internal/auth"
	"biliard_club/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerDeps struct {
	*config.JWTConfig
	*auth.Service
}
type Handler struct {
	*config.JWTConfig
	*auth.Service
}

func NewAuthHandler(engine *gin.Engine, deps HandlerDeps) {
	handler := Handler{
		Service:   deps.Service,
		JWTConfig: deps.JWTConfig,
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

func (h *Handler) Login(c *gin.Context) {
	var loginReq auth.LoginRequest
	if err := c.BindJSON(&loginReq); err != nil {
		return
	}
	user, err := h.Service.Login(loginReq.Phone, loginReq.Password)
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

func (h *Handler) Register(c *gin.Context) {
	var registerReq auth.RegisterRequest
	if err := c.BindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.Service.Register(registerReq.Phone, registerReq.Password, registerReq.Name)
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
