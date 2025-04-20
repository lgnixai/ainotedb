
package domain

import "time"

type Entity interface {
	GetID() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type BaseEntity struct {
	ID        string    `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e BaseEntity) GetID() string       { return e.ID }
func (e BaseEntity) GetCreatedAt() time.Time { return e.CreatedAt }
func (e BaseEntity) GetUpdatedAt() time.Time { return e.UpdatedAt }
