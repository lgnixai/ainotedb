package model

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// ViewType defines view type
type ViewType string

const (
	ViewTypeGrid     ViewType = "grid"
	ViewTypeGallery  ViewType = "gallery"
	ViewTypeKanban   ViewType = "kanban"
	ViewTypeCalendar ViewType = "calendar"
	ViewTypeTimeline ViewType = "timeline"

	ViewTypeChart ViewType = "chart"
)

// View represents a view
type View struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Type    ViewType     `json:"type"`
	TableID string       `json:"tableId"`
	Filter  *FilterGroup `json:"filter,omitempty"`
	Sort    []SortOption `json:"sort,omitempty"`
	Fields  []string     `json:"fields"` // Displayed field ID list
	Options ViewOptions  `json:"options,omitempty"`
	Config  string       `json:"config,omitempty"` //Added Config field
}

// ViewOptions represents view options
type ViewOptions struct {
	// Grid view options
	FrozenColumns []string `json:"frozenColumns,omitempty"`
	RowHeight     int      `json:"rowHeight,omitempty"`

	// Kanban view options
	GroupField string `json:"groupField,omitempty"`

	// Calendar view options
	DateField string `json:"dateField,omitempty"`

	// Gallery view options
	CoverField string `json:"coverField,omitempty"`
	CardSize   string `json:"cardSize,omitempty"`
}

// SortOption represents sort options
type SortOption struct {
	FieldID   string `json:"fieldId"`
	Direction string `json:"direction"` // asc, desc
}

func GenerateID() string {
	return uuid.New().String()
}
func (v *View) Validate() error {
	if v.Name == "" {
		return ErrEmptyViewName
	}
	if v.TableID == "" {
		return ErrEmptyTableID
	}

	switch v.Type {
	case ViewTypeGrid, ViewTypeGallery, ViewTypeKanban, ViewTypeCalendar, ViewTypeTimeline:
		return nil
	default:
		return ErrInvalidViewType
	}
}

// UpdateConfig updates the view configuration
func (v *View) UpdateConfig(config interface{}) error {
	configBytes, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	v.Config = string(configBytes)
	return nil
}

//Error handling for validation

// FilterGroup struct
type FilterGroup struct {
	Filters []Filter `json:"filters"`
}

// Filter struct
type Filter struct {
	Field string `json:"field"`
	Op    string `json:"op"`
	Value string `json:"value"`
}

func (v *View) GetTableID() string {
	return v.TableID
}
