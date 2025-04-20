package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/user/model"
	"github.com/undb/undb-go/internal/user/service"
	"github.com/undb/undb-go/internal/user/util"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	service service.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Register handles user registration
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &model.User{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Name,
	}

	if err := h.service.Register(c.Request.Context(), user); err != nil {
		if err == model.ErrEmailAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login handles user login
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		if err == model.ErrUserNotFound || err == model.ErrInvalidPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 生成 JWT token
	token, err := util.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}
	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"token": token,
	})
}

// GetUser handles getting a user by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == model.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser handles updating a user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = id
	if err := h.service.Update(c.Request.Context(), &user); err != nil {
		if err == model.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser handles deleting a user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		if err == model.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
