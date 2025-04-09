package table

import (
	"biliard_club/internal/table"
	"biliard_club/pkg/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HandlerDeps struct {
	Service *table.Service
}

type Handler struct {
	*table.Service
}

func NewHandler(engine *gin.Engine, deps HandlerDeps) {
	handler := &Handler{deps.Service}

	protected := engine.Group("/table")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/create", handler.Create)
		protected.PATCH("/update", handler.Update)
		protected.GET("/:id", handler.GetAll)
		protected.GET("/", handler.GetAll)

	}
}

func (h *Handler) GetAll(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		tables, err := h.Service.GetAllTables()
		if err != nil && errors.Is(err, table.NoTables) {
			c.JSON(http.StatusNoContent, gin.H{"error": table.NoTables})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, tables)
		return
	}
	if id, err := strconv.Atoi(id); err == nil {
		t, err := h.Service.GetByID(uint(id))
		if errors.Is(err, table.NotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, t)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id should be int"})
	}
}

func (h *Handler) Create(c *gin.Context) {
	var t table.Table
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	} else if t.PriceAfterSwitch == 0 || t.PriceBeforeSwitch == 0 ||
		t.SwitchLong == 0 || t.SwitchTime == 0 || t.Type == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	_, err := h.Service.Create(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) Update(c *gin.Context) {
	var t table.Table
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if t.ID == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	err := h.Service.Update(&t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}
