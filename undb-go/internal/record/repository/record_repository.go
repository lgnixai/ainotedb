package repository

import (
	"context"
	"fmt"

	"github.com/undb/undb-go/internal/record/model"
	"gorm.io/gorm"
)

// RecordRepository defines the interface for record data access
type RecordRepository interface {
	Create(ctx context.Context, record *model.Record) error
	GetByID(ctx context.Context, id string) (*model.Record, error)
	GetByTableID(ctx context.Context, tableID string) ([]*model.Record, error)
	Update(ctx context.Context, record *model.Record) error
	Delete(ctx context.Context, id string) error
	BatchCreate(ctx context.Context, records []*model.Record) (int64, error)
	BatchUpdate(ctx context.Context, records []*model.Record) (int64, error)
	BatchDelete(ctx context.Context, ids []string) (int64, error)
	Aggregate(ctx context.Context, req model.AggregationRequest) (interface{}, error)
	Pivot(ctx context.Context, req model.PivotRequest) (interface{}, error)
}

type recordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{db: db}
}

func (r *recordRepository) Create(ctx context.Context, record *model.Record) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *recordRepository) GetByID(ctx context.Context, id string) (*model.Record, error) {
	var record model.Record
	err := r.db.WithContext(ctx).First(&record, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrRecordNotFound
		}
		return nil, err
	}
	return &record, nil
}

func (r *recordRepository) GetByTableID(ctx context.Context, tableID string) ([]*model.Record, error) {
	var records []*model.Record
	err := r.db.WithContext(ctx).Where("table_id = ?", tableID).Find(&records).Error
	return records, err
}

func (r *recordRepository) Update(ctx context.Context, record *model.Record) error {
	return r.db.WithContext(ctx).Save(record).Error
}

func (r *recordRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Record{}, "id = ?", id).Error
}

func (r *recordRepository) BatchCreate(ctx context.Context, records []*model.Record) (int64, error) {
	if len(records) == 0 {
		return 0, nil
	}
	result := r.db.WithContext(ctx).CreateInBatches(records, 100)
	return result.RowsAffected, result.Error
}

func (r *recordRepository) BatchUpdate(ctx context.Context, records []*model.Record) (int64, error) {
	if len(records) == 0 {
		return 0, nil
	}

	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	var count int64
	for _, record := range records {
		result := tx.Save(record)
		if result.Error != nil {
			tx.Rollback()
			return count, result.Error
		}
		count += result.RowsAffected
	}

	return count, tx.Commit().Error
}

func (r *recordRepository) BatchDelete(ctx context.Context, ids []string) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	result := r.db.WithContext(ctx).Delete(&model.Record{}, "id IN ?", ids)
	return result.RowsAffected, result.Error
}

func (r *recordRepository) Aggregate(ctx context.Context, req model.AggregationRequest) (interface{}, error) {
	query := r.db.WithContext(ctx).Model(&model.Record{}).Where("table_id = ?", req.TableID)

	if req.Filter != "" {
		query = query.Where(req.Filter)
	}

	var result interface{}
	switch req.Function {
	case "count":
		var count int64
		err := query.Count(&count).Error
		result = count
		return result, err
	case "sum":
		var sum float64
		err := query.Select(fmt.Sprintf("SUM(%s)", req.Field)).Scan(&sum).Error
		result = sum
		return result, err
	case "avg":
		var avg float64
		err := query.Select(fmt.Sprintf("AVG(%s)", req.Field)).Scan(&avg).Error
		result = avg
		return result, err
	case "min":
		var min float64
		err := query.Select(fmt.Sprintf("MIN(%s)", req.Field)).Scan(&min).Error
		result = min
		return result, err
	case "max":
		var max float64
		err := query.Select(fmt.Sprintf("MAX(%s)", req.Field)).Scan(&max).Error
		result = max
		return result, err
	default:
		return nil, fmt.Errorf("unsupported aggregation function: %s", req.Function)
	}
}

func (r *recordRepository) Pivot(ctx context.Context, req model.PivotRequest) (interface{}, error) {
	// Implement pivot table with raw SQL
	query := r.db.WithContext(ctx).Raw(
		`SELECT 
			ROW_DIMENSION,
			GROUP_CONCAT(CASE WHEN COL_DIMENSION = ? THEN MEASURE END) as PIVOT_DATA
		FROM records 
		WHERE table_id = ?
		GROUP BY ROW_DIMENSION`,
		req.ColDimension, req.TableID)

	var results []map[string]interface{}
	err := query.Scan(&results).Error
	return results, err
}
