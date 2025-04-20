package repository

import (
	"context"

	"github.com/undb/undb-go/internal/record/model"
	"gorm.io/gorm"
)

// RecordRepository 定义记录仓库接口
type RecordRepository interface {
	// Create 创建新记录
	Create(ctx context.Context, record *model.Record) error

	// GetByID 根据ID获取记录
	GetByID(ctx context.Context, id string) (*model.Record, error)

	// GetByTableID 获取表的所有记录
	GetByTableID(ctx context.Context, tableID string) ([]*model.Record, error)

	// Update 更新记录
	Update(ctx context.Context, record *model.Record) error

	// Delete 删除记录
	Delete(ctx context.Context, id string) error

+	// BatchCreate 批量创建记录
+	// 返回成功创建的数量和遇到的第一个错误
+	BatchCreate(ctx context.Context, records []*model.Record) (int64, error)
+
	// BatchUpdate 批量更新记录
	// 注意：此方法可能需要根据具体的数据库和 GORM 版本进行调整以获得最佳性能和原子性。
	// 一个简单的实现是逐条更新，但效率较低。
	// 更高效的方法可能涉及事务或特定数据库的批量更新语法。
	BatchUpdate(ctx context.Context, records []*model.Record) (int64, error)

+	// BatchDelete 批量删除记录
+	// 返回成功删除的数量和遇到的第一个错误
+	BatchDelete(ctx context.Context, ids []string) (int64, error)
+
+	// Aggregate 执行聚合查询
+	// 返回聚合结果和错误。结果的类型取决于查询。
+	Aggregate(ctx context.Context, req model.AggregationRequest) (interface{}, error)
+
+	// Pivot 执行透视表查询
+	// 返回透视表数据和错误。数据的结构取决于实现。
+	Pivot(ctx context.Context, req model.PivotRequest) (interface{}, error)
}

// recordRepository 实现记录仓库接口
type recordRepository struct {
	db *gorm.DB
}

// NewRecordRepository 创建新的记录仓库实例
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
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *recordRepository) Update(ctx context.Context, record *model.Record) error {
	return r.db.WithContext(ctx).Save(record).Error
}

func (r *recordRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Record{}, "id = ?", id).Error
}

+// BatchCreate 批量创建记录 (使用 GORM 的 CreateInBatches)
+func (r *recordRepository) BatchCreate(ctx context.Context, records []*model.Record) (int64, error) {
+	if len(records) == 0 {
+		return 0, nil
+	}
+	// GORM 的 CreateInBatches 会处理 BeforeCreate 钩子
+	result := r.db.WithContext(ctx).CreateInBatches(records, 100) // Adjust batch size as needed
+	return result.RowsAffected, result.Error
+}
+
// BatchUpdate 批量更新记录 (简单实现，逐条更新在事务中)
// 注意：对于大量记录，此方法效率不高。考虑使用数据库特定的批量操作。
func (r *recordRepository) BatchUpdate(ctx context.Context, records []*model.Record) (int64, error) {
	var updatedCount int64
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	for _, record := range records {
		// 使用 Save 来处理插入或更新，确保 BeforeUpdate 钩子被调用
		result := tx.Save(record)
		if result.Error != nil {
			tx.Rollback()
			// 可以考虑收集所有错误而不是只返回第一个
			return updatedCount, result.Error
		}
		// 只有实际发生更新时才增加计数器 (GORM Save 会在未更改时不更新)
		if result.RowsAffected > 0 {
			updatedCount++
		}
	}

	if err := tx.Commit().Error; err != nil {
		return updatedCount, err
	}

	return updatedCount, nil
}

+// BatchDelete 批量删除记录
+func (r *recordRepository) BatchDelete(ctx context.Context, ids []string) (int64, error) {
+	if len(ids) == 0 {
+		return 0, nil
+	}
+	// 直接使用 Delete 方法，传入 ID 切片
+	result := r.db.WithContext(ctx).Where("id IN ?", ids).Delete(&model.Record{})
+	return result.RowsAffected, result.Error
+}
+
+// Aggregate 执行聚合查询 (Placeholder - Needs specific implementation)
+func (r *recordRepository) Aggregate(ctx context.Context, req model.AggregationRequest) (interface{}, error) {
+	// TODO: Implement actual aggregation logic using GORM or raw SQL.
+	// This will involve dynamically building the query based on req.
+	// Example (Conceptual - requires proper quoting and handling):
+	var result interface{}
+	query := r.db.WithContext(ctx).Model(&model.Record{}).Where("table_id = ?", req.TableID)
+
+	if req.Filter != "" {
+		// WARNING: Directly using req.Filter can lead to SQL injection if not sanitized.
+		// Use parameterized queries or a safe query builder.
+		query = query.Where(req.Filter)
+	}
+
+	aggregationFunc := ""
+	selectField := "*"
+	switch req.Aggregation {
+	case model.Count:
+		aggregationFunc = "COUNT"
+	case model.Sum:
+		aggregationFunc = "SUM"
+		selectField = req.Field
+	case model.Avg:
+		aggregationFunc = "AVG"
+		selectField = req.Field
+	case model.Min:
+		aggregationFunc = "MIN"
+		selectField = req.Field
+	case model.Max:
+		aggregationFunc = "MAX"
+		selectField = req.Field
+	default:
+		return nil, fmt.Errorf("unsupported aggregation: %s", req.Aggregation)
+	}
+
+	// Handle GroupBy
+	if len(req.GroupBy) > 0 {
+		// Need to select group keys and the aggregation result
+		// Example: SELECT group_field, COUNT(*) FROM records GROUP BY group_field
+		selectClause := strings.Join(req.GroupBy, ", ") + ", " + aggregationFunc + "(" + selectField + ") as result"
+		rows, err := query.Select(selectClause).Group(strings.Join(req.GroupBy, ", ")).Rows()
+		if err != nil {
+			return nil, err
+		}
+		defer rows.Close()
+		// TODO: Scan rows into a map or slice of structs
+		result = "Grouped aggregation result placeholder"
+	} else {
+		// Single aggregation result
+		// Example: SELECT COUNT(*) FROM records
+		selectClause := aggregationFunc + "(" + selectField + ")"
+		if err := query.Select(selectClause).Scan(&result).Error; err != nil {
+			return nil, err
+		}
+	}
+
+	log.Printf("Executed aggregation (placeholder): %+v, Result: %v", req, result)
+	return result, nil // Placeholder
+}
+
+// Pivot 执行透视表查询 (Placeholder - Needs specific implementation)
+func (r *recordRepository) Pivot(ctx context.Context, req model.PivotRequest) (interface{}, error) {
+	// TODO: Implement pivot table logic. This is highly database-specific.
+	// - PostgreSQL: Use CROSSTAB function (requires tablefunc extension).
+	// - MySQL: Use conditional aggregation (SUM(CASE WHEN ...)).
+	// - SQLite: Requires complex manual pivoting in Go after fetching data.
+	log.Printf("Executing pivot (placeholder): %+v", req)
+	return "Pivot table data placeholder", nil // Placeholder
+}
+
+// Add imports if not already present
+import (
+	"fmt"
+	"log"
+	"strings"
+
+	"gorm.io/gorm/clause"
+)
+
 // BatchUpdateOptimized 批量更新记录 (使用 GORM 的 Clauses ON CONFLICT/UPDATE)
 // 注意: 这依赖于数据库支持（如 PostgreSQL, MySQL 8.0.19+）。SQLite 可能需要不同的方法。
 // 并且需要确保 GORM 版本支持。
