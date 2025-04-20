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