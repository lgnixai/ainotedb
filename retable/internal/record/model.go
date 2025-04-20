package record

import (
	"time"
)

type Record struct {
	ID        string                 `json:"id" gorm:"primarykey"`
	TableID   string                 `json:"table_id"`
	Values    map[string]interface{} `json:"values" gorm:"type:jsonb"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

type RecordQuery struct {
	Filter  map[string]interface{} `json:"filter,omitempty"`
	Sort    []SortOption          `json:"sort,omitempty"`
	Page    int                   `json:"page,omitempty"`
	PerPage int                   `json:"per_page,omitempty"`
}

type SortOption struct {
	Field string `json:"field"`
	Order string `json:"order"` // asc or desc
}

type BulkOperation struct {
	Records []Record `json:"records"`
}