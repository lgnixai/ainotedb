package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// FieldType defines field types
type FieldType string

const (
	FieldTypeText        FieldType = "text"
	FieldTypeNumber      FieldType = "number"
	FieldTypeBoolean     FieldType = "boolean"
	FieldTypeDate        FieldType = "date"
	FieldTypeSelect      FieldType = "select"
	FieldTypeMultiSelect FieldType = "multi_select"
	FieldTypeReference   FieldType = "reference"
	FieldTypeLookup      FieldType = "lookup"
	FieldTypeRollup      FieldType = "rollup"
)

// FieldOptions represents field options
type FieldOptions struct {
	// Reference field options
	ForeignTableID   string `json:"foreignTableId,omitempty"`
	SymmetricFieldID string `json:"symmetricFieldId,omitempty"`

	// Lookup field options
	ReferenceFieldID string `json:"referenceFieldId,omitempty"`
	DisplayFieldID   string `json:"displayFieldId,omitempty"`

	// Rollup field options
	RollupFieldID     string `json:"rollupFieldId,omitempty"`
	AggregateFunction string `json:"fn,omitempty"` // count, sum, avg, etc.
}

// Field represents a field
type Field struct {
	ID          string       `json:"id" gorm:"primaryKey"`
	TableID     string       `json:"table_id" gorm:"index"`
	Name        string       `json:"name"`
	Type        FieldType    `json:"type"`
	Description string       `json:"description"`
	Required    bool         `json:"required"`
	Unique      bool         `json:"unique"`
	Options     FieldOptions `json:"options,omitempty" gorm:"type:text"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// TableName specifies table name
func (Field) TableName() string {
	return "fields"
}

// BeforeCreate before create hook
func (f *Field) BeforeCreate(tx *gorm.DB) error {
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate before update hook
func (f *Field) BeforeUpdate(tx *gorm.DB) error {
	f.UpdatedAt = time.Now()
	return nil
}

// Validate validates field data
func (f *Field) Validate() error {
	if f.Name == "" {
		return ErrEmptyFieldName
	}

	switch f.Type {
	case FieldTypeReference:
		if f.Options.ForeignTableID == "" {
			return ErrInvalidReferenceField
		}
	case FieldTypeLookup:
		if f.Options.ReferenceFieldID == "" || f.Options.DisplayFieldID == "" {
			return ErrInvalidLookupField
		}
	case FieldTypeRollup:
		if f.Options.RollupFieldID == "" || f.Options.ReferenceFieldID == "" {
			return ErrInvalidRollupField
		}
	}

	return nil
}

// isValidFieldType checks if field type is valid
func isValidFieldType(t FieldType) bool {
	switch t {
	case FieldTypeText, FieldTypeNumber, FieldTypeBoolean, FieldTypeDate,
		FieldTypeSelect, FieldTypeMultiSelect, FieldTypeReference, FieldTypeLookup, FieldTypeRollup:
		return true
	default:
		return false
	}
}

var (
	ErrInvalidReferenceField = &Error{Code: "invalid_reference_field", Message: "invalid reference field"}
	ErrInvalidLookupField    = &Error{Code: "invalid_lookup_field", Message: "invalid lookup field"}
	ErrInvalidRollupField    = &Error{Code: "invalid_rollup_field", Message: "invalid rollup field"}
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
