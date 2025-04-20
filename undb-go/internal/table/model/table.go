package model

import (
	"time"

	"gorm.io/gorm"
)

// Table 表示一个表格
type Table struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SpaceID     uint      `json:"space_id" gorm:"index"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Table) TableName() string {
	return "tables"
}

// BeforeCreate 创建前的钩子
func (t *Table) BeforeCreate(tx *gorm.DB) error {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前的钩子
func (t *Table) BeforeUpdate(tx *gorm.DB) error {
	t.UpdatedAt = time.Now()
	return nil
}

// Validate 验证表格数据
func (t *Table) Validate() error {
	if t.Name == "" {
		return ErrEmptyTableName
	}
	if t.SpaceID == 0 {
		return ErrEmptySpaceID
	}
	return nil
}
