package table

import (
	"time"
)

type Table struct {
	ID          string    `json:"id" gorm:"primarykey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SpaceID     string    `json:"space_id"`
	Schema      []Field   `json:"schema" gorm:"foreignKey:TableID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Field struct {
	ID          string    `json:"id" gorm:"primarykey"`
	TableID     string    `json:"table_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Required    bool      `json:"required"`
	Unique      bool      `json:"unique"`
	Options     []Option  `json:"options,omitempty" gorm:"foreignKey:FieldID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Option struct {
	ID      string `json:"id" gorm:"primarykey"`
	FieldID string `json:"field_id"`
	Label   string `json:"label"`
	Value   string `json:"value"`
}