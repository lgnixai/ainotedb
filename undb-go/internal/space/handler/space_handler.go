package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/space/model"
	"github.com/undb/undb-go/internal/space/service"
)

// SpaceHandler 空间处理器
type SpaceHandler struct {
	service service.SpaceService
}

// NewSpaceHandler 创建空间处理器实例
func NewSpaceHandler(service service.SpaceService) *SpaceHandler {
	return &SpaceHandler{service: service}
}

// CreateSpace 创建空间
func (h *SpaceHandler) CreateSpace(c *gin.Context) {
	var space model.Space
	if err := c.ShouldBindJSON(&space); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(c.Request.Context(), &space); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, space)
}

// GetSpace 获取空间
func (h *SpaceHandler) GetSpace(c *gin.Context) {
	id := c.Param("id")
	space, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, space)
}

// ListSpaces 获取空间列表
func (h *SpaceHandler) ListSpaces(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	spaces, err := h.service.GetByOwnerID(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"spaces": spaces})
}

// UpdateSpace 更新空间
func (h *SpaceHandler) UpdateSpace(c *gin.Context) {
	id := c.Param("id")
	var space model.Space
	if err := c.ShouldBindJSON(&space); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	space.ID = id
	if err := h.service.Update(c.Request.Context(), &space); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, space)
}

// DeleteSpace 删除空间
func (h *SpaceHandler) DeleteSpace(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateSpaceRequest represents the request body for creating a space
type CreateSpaceRequest struct {
	Name        string                `json:"name" binding:"required"`
	Description string                `json:"description"`
	Visibility  model.SpaceVisibility `json:"visibility" binding:"required"`
}

// Create handles the creation of a new space
func (h *SpaceHandler) Create(c *gin.Context) {
	var req CreateSpaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	space := &model.Space{
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     userID.(string),
		Visibility:  req.Visibility,
	}

	if err := h.service.Create(c.Request.Context(), space); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, space)
}

// GetByOwnerID handles retrieving spaces by owner ID
func (h *SpaceHandler) GetByOwnerID(c *gin.Context) {
	ownerID := c.Param("owner_id")
	spaces, err := h.service.GetByOwnerID(c.Request.Context(), ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, spaces)
}

// UpdateSpaceRequest represents the request body for updating a space
type UpdateSpaceRequest struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Visibility  model.SpaceVisibility `json:"visibility"`
}

// Update handles updating a space
func (h *SpaceHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req UpdateSpaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	space, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "space not found"})
		return
	}

	// 更新字段
	if req.Name != "" {
		space.Name = req.Name
	}
	if req.Description != "" {
		space.Description = req.Description
	}
	if req.Visibility != "" {
		space.Visibility = req.Visibility
	}

	if err := h.service.Update(c.Request.Context(), space); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, space)
}

// Delete handles deleting a space
func (h *SpaceHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// AddMemberRequest represents the request body for adding a member
type AddMemberRequest struct {
	UserID string           `json:"user_id" binding:"required"`
	Role   model.MemberRole `json:"role" binding:"required"`
}

// AddMember handles adding a member to a space
func (h *SpaceHandler) AddMember(c *gin.Context) {
	spaceID := c.Param("id")
	var req AddMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AddMember(c.Request.Context(), spaceID, req.UserID, req.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// RemoveMember handles removing a member from a space
func (h *SpaceHandler) RemoveMember(c *gin.Context) {
	spaceID := c.Param("id")
	userID := c.Param("user_id")

	if err := h.service.RemoveMember(c.Request.Context(), spaceID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateMemberRoleRequest represents the request body for updating a member's role
type UpdateMemberRoleRequest struct {
	Role model.MemberRole `json:"role" binding:"required"`
}

// UpdateMemberRole handles updating a member's role
func (h *SpaceHandler) UpdateMemberRole(c *gin.Context) {
	spaceID := c.Param("id")
	userID := c.Param("user_id")
	var req UpdateMemberRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateMemberRole(c.Request.Context(), spaceID, userID, req.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// GetSpaceMembers handles retrieving all members of a space
func (h *SpaceHandler) GetSpaceMembers(c *gin.Context) {
	spaceID := c.Param("id")
	members, err := h.service.GetSpaceMembers(c.Request.Context(), spaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}
