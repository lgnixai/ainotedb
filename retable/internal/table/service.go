package table

import (
	"errors"
	"gorm.io/gorm"
)

type TableService struct {
	db *gorm.DB
}

func NewTableService(db *gorm.DB) *TableService {
	return &TableService{db: db}
}

func (s *TableService) CreateTable(table *Table) error {
	return s.db.Create(table).Error
}

func (s *TableService) GetTable(id string) (*Table, error) {
	var table Table
	if err := s.db.Preload("Schema").Preload("Schema.Options").First(&table, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &table, nil
}

func (s *TableService) ListTables(spaceID string) ([]Table, error) {
	var tables []Table
	if err := s.db.Where("space_id = ?", spaceID).Find(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}

func (s *TableService) UpdateTable(table *Table) error {
	return s.db.Save(table).Error
}

func (s *TableService) DeleteTable(id string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Delete fields and their options
		if err := tx.Where("table_id = ?", id).Delete(&Field{}).Error; err != nil {
			return err
		}
		
		// Delete the table
		if err := tx.Delete(&Table{}, "id = ?", id).Error; err != nil {
			return err
		}
		
		return nil
	})
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