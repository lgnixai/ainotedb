package model

import (
	"database/sql/driver"
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
	Sort    SortOptionList `json:"sort,omitempty"`
	Fields  StringSlice  `json:"fields"` // Displayed field ID list
	Options ViewOptions  `json:"options,omitempty"`
	Config  string       `json:"config,omitempty"` //Added Config field
}

// StringSlice 用于 GORM JSON序列化

type StringSlice []string

func (ss StringSlice) Value() (driver.Value, error) {
	return json.Marshal(ss)
}

func (ss *StringSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal StringSlice value: %v", value)
	}
	return json.Unmarshal(bytes, ss)
}

// SortOptionList 用于 GORM JSON序列化

type SortOptionList []SortOption

func (sl SortOptionList) Value() (driver.Value, error) {
	return json.Marshal(sl)
}

func (sl *SortOptionList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal SortOptionList value: %v", value)
	}
	return json.Unmarshal(bytes, sl)
}


// ViewOptions represents view options

func (vo ViewOptions) Value() (driver.Value, error) {
	return json.Marshal(vo)
}

func (vo *ViewOptions) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal ViewOptions value: %v", value)
	}
	return json.Unmarshal(bytes, vo)
}

type ViewOptions struct {
	// Grid view options
	FrozenColumns StringSlice `json:"frozenColumns,omitempty"`
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
	Filters FilterList `json:"filters"`
}

// GORM JSON序列化支持
func (fg FilterGroup) Value() (driver.Value, error) {
	return json.Marshal(fg)
}

func (fg *FilterGroup) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal FilterGroup value: %v", value)
	}
	return json.Unmarshal(bytes, fg)
}

type FilterList []Filter

// GORM JSON序列化支持
func (fl FilterList) Value() (driver.Value, error) {
	return json.Marshal(fl)
}

func (fl *FilterList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal FilterList value: %v", value)
	}
	return json.Unmarshal(bytes, fl)
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
