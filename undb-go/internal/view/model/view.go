package model

import (
	"time"

	"gorm.io/gorm"
)

// ViewType represents the type of view
type ViewType string

const (
	ViewTypeGrid     ViewType = "grid"
	ViewTypeForm     ViewType = "form"
	ViewTypeGallery  ViewType = "gallery"
	ViewTypeKanban   ViewType = "kanban"
	ViewTypeGantt    ViewType = "gantt"
	ViewTypeCalendar ViewType = "calendar"
)

// View represents a view configuration
type View struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	TableID     string    `gorm:"index" json:"tableId"`
	Name        string    `json:"name"`
	Type        ViewType  `json:"type"`
	Description string    `json:"description"`
	Config      string    `gorm:"type:json" json:"config"` // JSON configuration for view-specific settings
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TableName specifies the table name for View
func (View) TableName() string {
	return "views"
}

// BeforeCreate is a GORM hook that runs before creating a view
func (v *View) BeforeCreate(tx *gorm.DB) error {
	if v.ID == "" {
		v.ID = GenerateID("view")
	}
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate is a GORM hook that runs before updating a view
func (v *View) BeforeUpdate(tx *gorm.DB) error {
	v.UpdatedAt = time.Now()
	return nil
}

// Validate validates the view data
func (v *View) Validate() error {
	if v.Name == "" {
		return ErrEmptyViewName
	}
	if v.TableID == "" {
		return ErrEmptyTableID
	}
	if v.Type == "" {
		return ErrEmptyViewType
	}
	return nil
}
