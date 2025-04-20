package table

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TableService struct {
	db *gorm.DB
}

func NewTableService(db *gorm.DB) *TableService {
	return &TableService{db: db}
}

func (s *TableService) CreateTable(ctx context.Context, table *Table) error {
	if table.ID == "" {
		table.ID = uuid.New().String()
	}

	return s.db.WithContext(ctx).Create(table).Error
}

func (s *TableService) GetTable(ctx context.Context, id string) (*Table, error) {
	var table Table
	if err := s.db.WithContext(ctx).Preload("Schema").First(&table, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &table, nil
}

func (s *TableService) ListTables(ctx context.Context, spaceID string) ([]Table, error) {
	var tables []Table
	if err := s.db.WithContext(ctx).Where("space_id = ?", spaceID).Find(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}

func (s *TableService) UpdateTable(ctx context.Context, table *Table) error {
	if table.ID == "" {
		return errors.New("table ID is required")
	}

	return s.db.WithContext(ctx).Save(table).Error
}

func (s *TableService) DeleteTable(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Delete fields
		if err := tx.WithContext(ctx).Where("table_id = ?", id).Delete(&Field{}).Error; err != nil {
			return err
		}

		// Delete table
		if err := tx.WithContext(ctx).Delete(&Table{}, "id = ?", id).Error; err != nil {
			return err
		}

		return nil
	})
}

// Field Management
func (s *TableService) CreateField(ctx context.Context, field *Field) error {
	if field.ID == "" {
		field.ID = uuid.New().String()
	}

	return s.db.WithContext(ctx).Create(field).Error
}

func (s *TableService) GetField(ctx context.Context, id string) (*Field, error) {
	var field Field
	if err := s.db.WithContext(ctx).First(&field, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &field, nil
}

func (s *TableService) UpdateField(ctx context.Context, field *Field) error {
	if field.ID == "" {
		return errors.New("field ID is required")
	}

	return s.db.WithContext(ctx).Save(field).Error
}

func (s *TableService) DeleteField(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Delete(&Field{}, "id = ?", id).Error
}

func (s *TableService) ListFields(ctx context.Context, tableID string) ([]Field, error) {
	var fields []Field
	if err := s.db.WithContext(ctx).Where("table_id = ?", tableID).Find(&fields).Error; err != nil {
		return nil, err
	}
	return fields, nil
}

func (s *TableService) ValidateField(field *Field) error {
	if field.Name == "" {
		return errors.New("field name is required")
	}
	if field.Type == "" {
		return errors.New("field type is required")
	}
	return nil
}

func (s *TableService) AddField(field *Field) error {
	// Validate field type
	if !isValidFieldType(field.Type) {
		return errors.New("invalid field type")
	}

	// Check if field name exists
	var count int64
	if err := s.db.Model(&Field{}).Where("table_id = ? AND name = ?", field.TableID, field.Name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("field name already exists")
	}

	return s.db.Create(field).Error
}

func (s *TableService) UpdateFieldOptions(fieldID string, options []Option) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Delete existing options
		if err := tx.Where("field_id = ?", fieldID).Delete(&Option{}).Error; err != nil {
			return err
		}
		// Create new options
		for _, opt := range options {
			opt.FieldID = fieldID
			if err := tx.Create(&opt).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func isValidFieldType(fieldType string) bool {
	validTypes := map[string]bool{
		"text":     true,
		"number":   true,
		"date":     true,
		"boolean":  true,
		"select":   true,
		"file":     true,
		"user":     true,
		"currency": true,
	}
	return validTypes[fieldType]
}

func (s *TableService) UpdateField(field *Field) error {
	return s.db.Save(field).Error
}

func (s *TableService) DeleteField(id string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Delete field options
		if err := tx.Where("field_id = ?", id).Delete(&Option{}).Error; err != nil {
			return err
		}

		// Delete the field
		if err := tx.Delete(&Field{}, "id = ?", id).Error; err != nil {
			return err
		}

		return nil
	})
}