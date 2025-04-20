package record

import (
	"errors"
	"gorm.io/gorm"
)

type RecordService struct {
	db *gorm.DB
}

func NewRecordService(db *gorm.DB) *RecordService {
	return &RecordService{db: db}
}

func (s *RecordService) ValidateRecord(tableID string, values map[string]interface{}) error {
	var fields []Field
	if err := s.db.Where("table_id = ?", tableID).Find(&fields).Error; err != nil {
		return err
	}

	for _, field := range fields {
		value, exists := values[field.ID]
		if field.Required && (!exists || value == nil) {
			return fmt.Errorf("field %s is required", field.Name)
		}
		if field.Unique && exists && value != nil {
			var count int64
			if err := s.db.Model(&Record{}).
				Where("table_id = ? AND values->>'$.?'= ?", tableID, field.ID, value).
				Count(&count).Error; err != nil {
				return err
			}
			if count > 0 {
				return fmt.Errorf("field %s must be unique", field.Name)
			}
		}
	}
	return nil
}

func (s *RecordService) BulkUpsertRecords(records []Record) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		for _, record := range records {
			if err := tx.Save(&record).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *RecordService) GetRecordsWithFilter(tableID string, filter map[string]interface{}, sort []string, page, pageSize int) ([]Record, int64, error) {
	var records []Record
	var total int64

	query := s.db.Model(&Record{}).Where("table_id = ?", tableID)

	// Apply filters
	for field, value := range filter {
		query = query.Where("values->>'$.?' = ?", field, value)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting and pagination
	if len(sort) > 0 {
		for _, s := range sort {
			query = query.Order(s)
		}
	}

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	if err := query.Find(&records).Error; err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

func (s *RecordService) CreateRecord(record *Record) error {
	return s.db.Create(record).Error
}

func (s *RecordService) BulkCreateRecords(records []Record) error {
	return s.db.Create(&records).Error
}

func (s *RecordService) GetRecord(id string) (*Record, error) {
	var record Record
	if err := s.db.First(&record, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (s *RecordService) ListRecords(tableID string, query *RecordQuery) ([]Record, error) {
	var records []Record
	db := s.db.Where("table_id = ?", tableID)

	if query != nil {
		if query.Filter != nil {
			db = db.Where(query.Filter)
		}

		if len(query.Sort) > 0 {
			for _, sort := range query.Sort {
				db = db.Order(sort.Field + " " + sort.Order)
			}
		}

		if query.Page > 0 && query.PerPage > 0 {
			offset := (query.Page - 1) * query.PerPage
			db = db.Offset(offset).Limit(query.PerPage)
		}
	}

	if err := db.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (s *RecordService) UpdateRecord(record *Record) error {
	return s.db.Save(record).Error
}

func (s *RecordService) BulkUpdateRecords(records []Record) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		for _, record := range records {
			if err := tx.Save(&record).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *RecordService) DeleteRecord(id string) error {
	return s.db.Delete(&Record{}, "id = ?", id).Error
}

func (s *RecordService) BulkDeleteRecords(ids []string) error {
	return s.db.Delete(&Record{}, "id IN ?", ids).Error
}