package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/table/model"
	"github.com/undb/undb-go/internal/table/service"
)

// TableHandler handles HTTP requests for tables
type TableHandler struct {
	service service.TableService
}

// NewTableHandler creates a new TableHandler instance
func NewTableHandler(service service.TableService) *TableHandler {
	return &TableHandler{service: service}
}

// CreateTableRequest represents the request body for creating a table
type CreateTableRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	SpaceID     string `json:"space_id" binding:"required"`
}

// CreateTable 创建新表
func (h *TableHandler) CreateTable(c *gin.Context) {
	var table model.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(c.Request.Context(), &table); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, table)
}

// GetTable 获取表
func (h *TableHandler) GetTable(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	id := uint(id64)
	table, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, table)
}

// GetTables 获取空间的所有表
func (h *TableHandler) GetTables(c *gin.Context) {
	spaceIDStr := c.Param("space_id")
	spaceID64, err := strconv.ParseUint(spaceIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid space_id"})
		return
	}
	spaceID := uint(spaceID64)
	tables, err := h.service.GetBySpaceID(c.Request.Context(), spaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tables)
}

// GetTablesBySpace 获取空间下的所有表格
func (h *TableHandler) GetTablesBySpace(c *gin.Context) {
	spaceIDStr := c.Param("space_id")
	spaceID64, err := strconv.ParseUint(spaceIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid space_id"})
		return
	}
	spaceID := uint(spaceID64)
	if spaceID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "space_id is required"})
		return
	}

	tables, err := h.service.GetBySpaceID(c.Request.Context(), spaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tables)
}

// UpdateTableRequest represents the request body for updating a table
type UpdateTableRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateTable 更新表
func (h *TableHandler) UpdateTable(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	id := uint(id64)
	var table model.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	table.ID = id
	if err := h.service.Update(c.Request.Context(), &table); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, table)
}

// DeleteTable 删除表
func (h *TableHandler) DeleteTable(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	id := uint(id64)
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
