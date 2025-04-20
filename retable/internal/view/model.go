package view

import (
	"time"
)

type ViewType string

const (
	GridView    ViewType = "grid"
	KanbanView  ViewType = "kanban"
	GalleryView ViewType = "gallery"
	CalendarView ViewType = "calendar"
)

type View struct {
	ID          string    `json:"id"`
	TableID     string    `json:"table_id"`
	Name        string    `json:"name"`
	Type        ViewType  `json:"type"`
	Fields      []string  `json:"fields"`
	Filter      *Filter   `json:"filter,omitempty"`
	Sort        []Sort    `json:"sort,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Filter struct {
	Operator string        `json:"operator"`
	Rules    []FilterRule  `json:"rules"`
}

type FilterRule struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type Sort struct {
	Field string `json:"field"`
	Order string `json:"order"`
}