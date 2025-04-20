package domain

import (
	"time"
)

type Table struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"` 
	SpaceID   string    `json:"spaceId"`
	Schema    Schema    `json:"schema"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Schema struct {
	Fields []Field `json:"fields"`
}

type Field struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Required    bool        `json:"required"`
	Unique      bool        `json:"unique"`
	Options     interface{} `json:"options,omitempty"`
}

func (t *Table) GetID() string {
	return t.ID
}

func NewTable(id, name, spaceId string) *Table {
	return &Table{
		ID:        id,
		Name:      name,
		SpaceID:   spaceId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
