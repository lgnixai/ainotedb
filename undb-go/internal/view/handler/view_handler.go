package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/view/model"
	"github.com/undb/undb-go/internal/view/service"
)

// ViewHandler handles HTTP requests for views
type ViewHandler struct {
	service service.ViewService
}

// NewViewHandler creates a new ViewHandler
func NewViewHandler(service service.ViewService) *ViewHandler {
	return &ViewHandler{service: service}
}

// CreateView handles view creation
func (h *ViewHandler) CreateView(c *gin.Context) {
	var view model.View
	if err := c.ShouldBindJSON(&view); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateView(c.Request.Context(), &view); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, view)
}

// GetView handles getting a view by ID
func (h *ViewHandler) GetView(c *gin.Context) {
	id := c.Param("id")
	view, err := h.service.GetView(c.Request.Context(), id)
	if err != nil {
		if err == model.ErrViewNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, view)
}

// GetViews handles getting all views for a table
func (h *ViewHandler) GetViews(c *gin.Context) {
	tableID := c.Param("tableId")
	views, err := h.service.GetViews(c.Request.Context(), tableID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, views)
}

// UpdateView handles updating a view
func (h *ViewHandler) UpdateView(c *gin.Context) {
	id := c.Param("id")
	var view model.View
	if err := c.ShouldBindJSON(&view); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	view.ID = id
	if err := h.service.UpdateView(c.Request.Context(), &view); err != nil {
		if err == model.ErrViewNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, view)
}

// DeleteView handles deleting a view
func (h *ViewHandler) DeleteView(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteView(c.Request.Context(), id); err != nil {
		if err == model.ErrViewNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateViewConfig handles updating a view's configuration
func (h *ViewHandler) UpdateViewConfig(c *gin.Context) {
	id := c.Param("id")
	var config interface{}
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateViewConfig(c.Request.Context(), id, config); err != nil {
		if err == model.ErrViewNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
