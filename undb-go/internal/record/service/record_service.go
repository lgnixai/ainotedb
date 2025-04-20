package service

import (
	"context"
	"fmt"
	"log"
	"sync"
	"golang.org/x/sync/errgroup"

	"github.com/undb/undb-go/internal/infrastructure/db"
	"github.com/undb/undb-go/internal/record/model"
)

// RecordService 定义记录服务接口
type RecordService interface {
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

	// BatchCreateRecords 批量创建记录
	BatchCreateRecords(ctx context.Context, req model.BatchCreateRecordRequest) (*model.BatchCreateRecordResponse, error)

	// BatchUpdateRecords 批量更新记录
	BatchUpdateRecords(ctx context.Context, req model.BatchUpdateRecordRequest) (*model.BatchUpdateRecordResponse, error)

	// BatchDeleteRecords 批量删除记录
	BatchDeleteRecords(ctx context.Context, req model.BatchDeleteRecordRequest) (*model.BatchDeleteRecordResponse, error)

	// AggregateRecords 执行聚合查询
	AggregateRecords(ctx context.Context, req model.AggregationRequest) (*model.AggregationResponse, error)

	// PivotRecords 执行透视表查询
	PivotRecords(ctx context.Context, req model.PivotRequest) (*model.PivotResponse, error)
}

// recordService 实现记录服务接口
type recordService struct {
	repo repository.RecordRepository
}

// NewRecordService 创建新的记录服务实例
func NewRecordService(repo repository.RecordRepository) RecordService {
	return &recordService{repo: repo}
}

func (s *recordService) Create(ctx context.Context, record *model.Record) error {
	if record.ID == 0 {
		 // Assign UUID if not provided
	}
	if err := record.Validate(); err != nil {
		return err
	}
	return s.repo.Create(ctx, record)
}

func (s *recordService) GetByID(ctx context.Context, id string) (*model.Record, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *recordService) GetByTableID(ctx context.Context, tableID string) ([]*model.Record, error) {
	return s.repo.GetByTableID(ctx, tableID)
}

func (s *recordService) Update(ctx context.Context, record *model.Record) error {
	if err := record.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ctx, record)
}

func (s *recordService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// BatchCreateRecords 批量创建记录
func (s *recordService) BatchCreateRecords(ctx context.Context, req model.BatchCreateRecordRequest) (*model.BatchCreateRecordResponse, error) {
	var (
		successCount int
		failedCount  int
		createdIDs   []uint
		errors       []string
		mu           sync.Mutex
		eg, childCtx = errgroup.WithContext(ctx)
	)

	eg.SetLimit(10) // Limit concurrency

	for _, recordData := range req.Records {
		data := recordData // Capture loop variable
		eg.Go(func() error {
			record := &model.Record{
				TableID: req.TableID,
				Data:    data,
			}

			if err := record.Validate(); err != nil {
				mu.Lock()
				failedCount++
				errors = append(errors, fmt.Sprintf("Validation failed for record data %v: %v", data, err))
				mu.Unlock()
				return nil // Continue processing others
			}

			if err := s.repo.Create(childCtx, record); err != nil {
				mu.Lock()
				failedCount++
				errors = append(errors, fmt.Sprintf("Failed to create record data %v: %v", data, err))
				mu.Unlock()
				return nil // Continue processing others
			}

			mu.Lock()
			successCount++
			createdIDs = append(createdIDs, record.ID)
			mu.Unlock()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		// Log or handle the context error if necessary
		log.Printf("Error during batch create: %v", err)
		// Depending on requirements, you might want to return this error
	}

	return &model.BatchCreateRecordResponse{
		SuccessCount: successCount,
		FailedCount:  failedCount,
		CreatedIDs:   createdIDs,
		Errors:       errors,
	}, nil
}

// BatchUpdateRecords 批量更新记录
func (s *recordService) BatchUpdateRecords(ctx context.Context, req model.BatchUpdateRecordRequest) (*model.BatchUpdateRecordResponse, error) {
	var (
		successCount int
		failedIDs    []string
		mu           sync.Mutex // Mutex to protect shared variables
		eg, childCtx = errgroup.WithContext(ctx)
	)

	// Limit concurrency to avoid overwhelming the database
	eg.SetLimit(10) // Example limit, adjust as needed

	for _, recordData := range req.Records {
		// Capture loop variable for goroutine
		data := recordData
		eg.Go(func() error {
			// Find the existing record first
			existingRecord, err := s.repo.GetByID(childCtx, data.ID)
			if err != nil {
				mu.Lock()
				failedIDs = append(failedIDs, data.ID)
				mu.Unlock()
				// Log the error but continue processing others
				log.Printf("Failed to get record %s for update: %v", data.ID, err)
				return nil // Don't fail the whole batch for one record fetch error
			}

			// Update the record data
			existingRecord.Data = data.Data

			// Validate before updating
			if err := existingRecord.Validate(); err != nil {
				mu.Lock()
				failedIDs = append(failedIDs, data.ID)
				mu.Unlock()
				log.Printf("Validation failed for record %s: %v", data.ID, err)
				return nil // Continue processing others
			}

			// Perform the update
			if err := s.repo.Update(childCtx, existingRecord); err != nil {
				mu.Lock()
				failedIDs = append(failedIDs, data.ID)
				mu.Unlock()
				log.Printf("Failed to update record %s: %v", data.ID, err)
				return nil // Continue processing others
			}

			mu.Lock()
			successCount++
			mu.Unlock()
			return nil
		})
	}

	// Wait for all goroutines to complete.
	// The first non-nil error returned by a goroutine will be returned by Wait().
	if err := eg.Wait(); err != nil {
		// This error comes from the errgroup context or a returned error from a goroutine
		// If goroutines return nil on failure, this might not capture all failures.
		// Consider more robust error handling if needed.
		log.Printf("Error during batch update: %v", err)
		// return nil, err // Or handle the error more specifically
	}

	return &model.BatchUpdateRecordResponse{
		SuccessCount: successCount,
		FailedIDs:    failedIDs,
	}, nil
}

// BatchDeleteRecords 批量删除记录
func (s *recordService) BatchDeleteRecords(ctx context.Context, req model.BatchDeleteRecordRequest) (*model.BatchDeleteRecordResponse, error) {
	var (
		successCount int
		failedIDs    []string
		mu           sync.Mutex
		eg, childCtx = errgroup.WithContext(ctx)
	)

	eg.SetLimit(10) // Limit concurrency

	for _, recordID := range req.RecordIDs {
		id := recordID // Capture loop variable
		eg.Go(func() error {
			if err := s.repo.Delete(childCtx, id); err != nil {
				mu.Lock()
				failedIDs = append(failedIDs, id)
				mu.Unlock()
				log.Printf("Failed to delete record %s: %v", id, err)
				// Decide if a single failure should stop the batch or just be recorded
				return nil // Continue processing others
			}
			mu.Lock()
			successCount++
			mu.Unlock()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		log.Printf("Error during batch delete: %v", err)
		// Handle context error if necessary
	}

	return &model.BatchDeleteRecordResponse{
		SuccessCount: successCount,
		FailedIDs:    failedIDs,
	}, nil
}

// AggregateRecords 执行聚合查询
func (s *recordService) AggregateRecords(ctx context.Context, req model.AggregationRequest) (*model.AggregationResponse, error) {
	// Basic validation
	switch req.Aggregation {
	case model.Count:
		// Field is optional for COUNT(*)
	case model.Sum, model.Avg, model.Min, model.Max:
		if req.Field == "" {
			return nil, fmt.Errorf("field is required for aggregation type %s", req.Aggregation)
		}
	default:
		return nil, fmt.Errorf("unsupported aggregation type: %s", req.Aggregation)
	}

	// TODO: Implement the actual aggregation logic in the repository
	// This might involve constructing dynamic SQL queries based on the request.
	// For now, returning a placeholder.
	log.Printf("Aggregation requested: %+v (Repository implementation pending)", req)

	// Example call to a potential repository method (needs implementation)
	// result, err := s.repo.Aggregate(ctx, req)
	// if err != nil {
	// 	 return nil, fmt.Errorf("repository aggregation failed: %w", err)
	// }

	// Placeholder response
	var result interface{}
	if len(req.GroupBy) > 0 {
		result = make(map[string]interface{}) // Placeholder for grouped results
	} else {
		result = 0 // Placeholder for single result
	}

	return &model.AggregationResponse{Result: result}, nil
}

// PivotRecords 执行透视表查询
func (s *recordService) PivotRecords(ctx context.Context, req model.PivotRequest) (*model.PivotResponse, error) {
	// Basic validation
	if len(req.Rows) == 0 || len(req.Columns) == 0 || req.Values == "" || req.AggFunc == "" {
		return nil, fmt.Errorf("rows, columns, values, and agg_func are required for pivot table")
	}

	// TODO: Implement the actual pivot table logic in the repository.
	// This is complex and often requires dynamic SQL (e.g., using CROSSTAB in PostgreSQL)
	// or significant data manipulation in Go after fetching raw data.
	log.Printf("Pivot table requested: %+v (Repository implementation pending)", req)

	// Example call to a potential repository method (needs implementation)
	// data, err := s.repo.Pivot(ctx, req)
	// if err != nil {
	// 	 return nil, fmt.Errorf("repository pivot failed: %w", err)
	// }

	// Placeholder response
	return &model.PivotResponse{Data: "Pivot table data placeholder"}, nil
}
