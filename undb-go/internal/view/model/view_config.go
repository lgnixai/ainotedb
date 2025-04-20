package model

import "encoding/json"

// ViewConfig represents the base configuration for all view types
type ViewConfig struct {
	Type ViewType `json:"type"`
}

// GridViewConfig represents the configuration for grid view
type GridViewConfig struct {
	ViewConfig
	Sort   []SortConfig   `json:"sort"`
	Filter []FilterConfig `json:"filter"`
	Fields []string       `json:"fields"` // Field IDs to display
}

// KanbanViewConfig represents the configuration for kanban view
type KanbanViewConfig struct {
	ViewConfig
	GroupBy string         `json:"groupBy"` // Field ID to group by
	Sort    []SortConfig   `json:"sort"`
	Filter  []FilterConfig `json:"filter"`
}

// GalleryViewConfig represents the configuration for gallery view
type GalleryViewConfig struct {
	ViewConfig
	CoverField string         `json:"coverField"` // Field ID for cover image
	Sort       []SortConfig   `json:"sort"`
	Filter     []FilterConfig `json:"filter"`
}

// SortConfig represents the sorting configuration
type SortConfig struct {
	Field string `json:"field"` // Field ID
	Order string `json:"order"` // "asc" or "desc"
}

// FilterConfig represents the filtering configuration
type FilterConfig struct {
	Field    string      `json:"field"`    // Field ID
	Operator string      `json:"operator"` // "eq", "neq", "gt", "lt", etc.
	Value    interface{} `json:"value"`
}

// GetConfig returns the appropriate view configuration based on the view type
func (v *View) GetConfig() (interface{}, error) {
	switch v.Type {
	case ViewTypeGrid:
		var config GridViewConfig
		if err := json.Unmarshal([]byte(v.Config), &config); err != nil {
			return nil, err
		}
		return config, nil
	case ViewTypeKanban:
		var config KanbanViewConfig
		if err := json.Unmarshal([]byte(v.Config), &config); err != nil {
			return nil, err
		}
		return config, nil
	case ViewTypeGallery:
		var config GalleryViewConfig
		if err := json.Unmarshal([]byte(v.Config), &config); err != nil {
			return nil, err
		}
		return config, nil
	default:
		return nil, ErrInvalidViewType
	}
}

// SetConfig sets the view configuration based on the view type
func (v *View) SetConfig(config interface{}) error {
	configBytes, err := json.Marshal(config)
	if err != nil {
		return err
	}
	v.Config = string(configBytes)
	return nil
}
