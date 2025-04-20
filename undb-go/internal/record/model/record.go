package model

import (
	"time"

	"gorm.io/gorm"
)

// Record 表示一条记录
type Record struct {
	ID        string                 `json:"id" gorm:"primaryKey"`
	TableID   string                 `json:"table_id" gorm:"index"`
	Data      map[string]interface{} `json:"data"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

// TableName 指定表名
func (Record) TableName() string {
	return "records"
}

// BeforeCreate 创建前的钩子
func (r *Record) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	r.CreatedAt = now
	r.UpdatedAt = now
	return nil
}

// BeforeUpdate 更新前的钩子
func (r *Record) BeforeUpdate(tx *gorm.DB) error {
	r.UpdatedAt = time.Now()
	return nil
}

// Validate 验证记录数据
func (r *Record) Validate() error {
	if r.TableID == "" {
		return ErrEmptyTableID
	}
	// Allow empty data for updates, but maybe not for creates?
	// if len(r.Data) == 0 {
	// 	return ErrEmptyFields
	// }
	return nil
}

// --- Batch Operations ---

// BatchCreateRecordRequest 定义批量创建记录的请求结构
type BatchCreateRecordRequest struct {
	TableID string                   `json:"table_id" binding:"required"`
	Records []map[string]interface{} `json:"records" binding:"required,min=1"` // Array of record data objects
}

// BatchCreateRecordResponse 定义批量创建记录的响应结构
type BatchCreateRecordResponse struct {
	SuccessCount int      `json:"success_count"`
	FailedCount  int      `json:"failed_count"`
	CreatedIDs   []string `json:"created_ids"` // IDs of successfully created records
	Errors       []string `json:"errors,omitempty"` // Optional: Detailed errors for failed creations
}

// BatchUpdateRecordData 定义批量更新中单个记录的数据结构
type BatchUpdateRecordData struct {
	ID   string                 `json:"id" binding:"required"`
	Data map[string]interface{} `json:"data" binding:"required"`
}

// BatchUpdateRecordRequest 定义批量更新记录的请求结构
type BatchUpdateRecordRequest struct {
	Records []BatchUpdateRecordData `json:"records" binding:"required,min=1"`
}

// BatchUpdateRecordResponse 定义批量更新记录的响应结构
type BatchUpdateRecordResponse struct {
	SuccessCount int      `json:"success_count"`
	FailedIDs    []string `json:"failed_ids"`
}

// BatchDeleteRecordRequest 定义批量删除记录的请求结构
type BatchDeleteRecordRequest struct {
	RecordIDs []string `json:"record_ids" binding:"required,min=1"`
}

// BatchDeleteRecordResponse 定义批量删除记录的响应结构
type BatchDeleteRecordResponse struct {
	SuccessCount int      `json:"success_count"`
	FailedIDs    []string `json:"failed_ids"`
}

// --- Aggregation & Pivot --- (Placeholders - Define specific structures as needed)

// AggregationType 定义聚合类型
type AggregationType string

const (
	Count AggregationType = "COUNT"
	Sum   AggregationType = "SUM"
	Avg   AggregationType = "AVG"
	Min   AggregationType = "MIN"
	Max   AggregationType = "MAX"
)

// AggregationRequest 定义聚合查询请求
type AggregationRequest struct {
	TableID     string          `json:"table_id" binding:"required"`
	Aggregation AggregationType `json:"aggregation" binding:"required"`
	Field       string          `json:"field,omitempty"` // Required for SUM, AVG, MIN, MAX
	Filter      string          `json:"filter,omitempty"` // Optional filter criteria (e.g., SQL WHERE clause)
	GroupBy     []string        `json:"group_by,omitempty"` // Optional grouping fields
}

// AggregationResponse 定义聚合查询响应
type AggregationResponse struct {
	Result interface{} `json:"result"` // Can be a single value or a map/slice for grouped results
}

// PivotRequest 定义透视表查询请求
type PivotRequest struct {
	TableID string   `json:"table_id" binding:"required"`
	Rows    []string `json:"rows" binding:"required,min=1"`     // Fields for rows
	Columns []string `json:"columns" binding:"required,min=1"` // Fields for columns
	Values  string   `json:"values" binding:"required"`       // Field to aggregate for cell values
	AggFunc string   `json:"agg_func" binding:"required"`    // Aggregation function (e.g., SUM, COUNT)
	Filter  string   `json:"filter,omitempty"`               // Optional filter criteria
}

// PivotResponse 定义透视表查询响应 (Structure depends heavily on implementation)
type PivotResponse struct {
	Data interface{} `json:"data"` // Typically a 2D array or map representing the pivot table
}

// Error types (Consider moving to a dedicated errors package)
var (
	ErrEmptyTableID = errors.New("table_id cannot be empty")
	ErrEmptyFields  = errors.New("record data cannot be empty")
)

// Add import for errors
import "errors"
