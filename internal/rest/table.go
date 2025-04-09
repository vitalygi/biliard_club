package rest

import (
	"biliard_club/domain"
	"biliard_club/internal/rest/middleware"
	"biliard_club/internal/service/table"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TableHandlerDeps struct {
	TableService *table.Service
}

type TableHandler struct {
	TableService *table.Service
}

func NewTableHandler(engine *gin.Engine, deps TableHandlerDeps) {
	handler := &TableHandler{deps.TableService}

	protected := engine.Group("/table")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/create", handler.Create)
		protected.PATCH("/update", handler.Update)
		protected.GET("/:id", handler.GetAll)
		protected.GET("/", handler.GetAll)

	}
}

func (h *TableHandler) GetAll(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		tables, err := h.TableService.GetAllTables()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, tables)
		return
	}
	if id, err := strconv.Atoi(id); err == nil {
		t, err := h.TableService.GetByID(uint(id))
		if errors.Is(err, domain.ErrTableNotFound) {
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

func (h *TableHandler) Create(c *gin.Context) {
	var t domain.Table
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	} else if t.PriceAfterSwitch == 0 || t.PriceBeforeSwitch == 0 ||
		t.SwitchLong == 0 || t.SwitchTime == 0 || t.Type == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	_, err := h.TableService.Create(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *TableHandler) Update(c *gin.Context) {
	var t domain.Table
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if t.ID == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	err := h.TableService.Update(&t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}
