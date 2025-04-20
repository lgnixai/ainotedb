package model

import (
	"time"

	"gorm.io/gorm"
)

// FieldType 定义字段类型
type FieldType string

const (
	FieldTypeText        FieldType = "text"
	FieldTypeNumber      FieldType = "number"
	FieldTypeBoolean     FieldType = "boolean"
	FieldTypeDate        FieldType = "date"
	FieldTypeSelect      FieldType = "select"
	FieldTypeMultiSelect FieldType = "multi_select"
	FieldTypeReference   FieldType = "reference"
)

// Field 表示一个字段
type Field struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	TableID     string    `json:"table_id" gorm:"index"`
	Name        string    `json:"name"`
	Type        FieldType `json:"type"`
	Description string    `json:"description"`
	Required    bool      `json:"required"`
	Unique      bool      `json:"unique"`
	Options     string    `json:"options" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Field) TableName() string {
	return "fields"
}

// BeforeCreate 创建前的钩子
func (f *Field) BeforeCreate(tx *gorm.DB) error {
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前的钩子
func (f *Field) BeforeUpdate(tx *gorm.DB) error {
	f.UpdatedAt = time.Now()
	return nil
}

// Validate 验证字段数据
func (f *Field) Validate() error {
	if f.Name == "" {
		return ErrEmptyFieldName
	}
	if f.Type == "" {
		return ErrEmptyFieldType
	}
	if !isValidFieldType(f.Type) {
		return ErrInvalidFieldType
	}
	return nil
}

// isValidFieldType 检查字段类型是否有效
func isValidFieldType(t FieldType) bool {
	switch t {
	case FieldTypeText, FieldTypeNumber, FieldTypeBoolean, FieldTypeDate,
		FieldTypeSelect, FieldTypeMultiSelect, FieldTypeReference:
		return true
	default:
		return false
	}
}
