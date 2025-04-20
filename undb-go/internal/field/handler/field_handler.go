package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/field/model"
	"github.com/undb/undb-go/internal/field/service"
)

type FieldHandler struct {
	service service.FieldService
}

func NewFieldHandler(service service.FieldService) *FieldHandler {
	return &FieldHandler{service: service}
}

// CreateField 创建新字段
func (h *FieldHandler) CreateField(c *gin.Context) {
	var field model.Field
	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(c.Request.Context(), &field); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, field)
}

// GetField 获取字段
func (h *FieldHandler) GetField(c *gin.Context) {
	id := c.Param("id")
	field, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, field)
}

// GetFields 获取表的所有字段
func (h *FieldHandler) GetFields(c *gin.Context) {
	tableID := c.Param("table_id")
	fields, err := h.service.GetByTableID(c.Request.Context(), tableID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fields)
}

// UpdateField 更新字段
func (h *FieldHandler) UpdateField(c *gin.Context) {
	id := c.Param("id")
	var field model.Field
	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	field.ID = id
	if err := h.service.Update(c.Request.Context(), &field); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, field)
}

// DeleteField 删除字段
func (h *FieldHandler) DeleteField(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
