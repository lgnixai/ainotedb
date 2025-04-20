package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/record/model"
	"github.com/undb/undb-go/internal/record/service"
)

// RecordHandler 结构体
type RecordHandler struct {
	RecordService service.RecordService
}

// NewRecordHandler 创建一个新的 RecordHandler
func NewRecordHandler(recordService service.RecordService) *RecordHandler {
	return &RecordHandler{
		RecordService: recordService,
	}
}

// CreateRecord 创建记录
func (h *RecordHandler) CreateRecord(c *gin.Context) {
	var record model.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Rename h.service to h.RecordService for consistency if needed
	if err := h.RecordService.Create(c.Request.Context(), &record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, record)
}

// GetRecord 获取记录
func (h *RecordHandler) GetRecord(c *gin.Context) {
	id := c.Param("id")
	record, err := h.RecordService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, record)
}

// GetRecords 获取表的所有记录 (Consider renaming or clarifying if this is different from GetRecordsByTable)
func (h *RecordHandler) GetRecords(c *gin.Context) {
	tableID := c.Param("table_id") // Assuming table_id is a path parameter here
	if tableID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "table_id path parameter is required"})
		return
	}
	records, err := h.RecordService.GetByTableID(c.Request.Context(), tableID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

// UpdateRecord 更新记录
func (h *RecordHandler) UpdateRecord(c *gin.Context) {
	id := c.Param("id")
	var record model.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the ID from the path is used, not from the body if present
	record.ID = id
	if err := h.RecordService.Update(c.Request.Context(), &record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, record)
}

// DeleteRecord 删除记录
func (h *RecordHandler) DeleteRecord(c *gin.Context) {
	id := c.Param("id")
	if err := h.RecordService.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetRecordsByTable 获取表格下的所有记录 (Seems redundant with GetRecords if using path param)
func (h *RecordHandler) GetRecordsByTable(c *gin.Context) {
	tableID := c.Param("table_id")
	if tableID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "table_id is required"})
		return
	}

	records, err := h.RecordService.GetByTableID(c.Request.Context(), tableID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

// --- Batch Operations Handlers ---

// BatchCreateRecords 批量创建记录
func (h *RecordHandler) BatchCreateRecords(c *gin.Context) {
	var req model.BatchCreateRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	resp, err := h.RecordService.BatchCreateRecords(c.Request.Context(), req)
	if err != nil {
		// Consider more specific error handling based on service error type
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to batch create records: " + err.Error()})
		return
	}

	// Determine appropriate status code based on results (e.g., 207 Multi-Status if partial success)
	statusCode := http.StatusCreated
	if resp.FailedCount > 0 {
		statusCode = http.StatusMultiStatus
	}
	c.JSON(statusCode, resp)
}

// BatchUpdateRecords 批量更新记录
func (h *RecordHandler) BatchUpdateRecords(c *gin.Context) {
	var req model.BatchUpdateRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	resp, err := h.RecordService.BatchUpdateRecords(c.Request.Context(), req)
	if err != nil {
		// Consider more specific error handling based on service error type
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to batch update records: " + err.Error()})
		return
	}

	// Determine appropriate status code (e.g., 207 Multi-Status if partial success)
	statusCode := http.StatusOK
	if len(resp.FailedIDs) > 0 {
		statusCode = http.StatusMultiStatus
	}
	c.JSON(statusCode, resp)
}

// BatchDeleteRecords 批量删除记录
func (h *RecordHandler) BatchDeleteRecords(c *gin.Context) {
	var req model.BatchDeleteRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	resp, err := h.RecordService.BatchDeleteRecords(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to batch delete records: " + err.Error()})
		return
	}

	// Determine appropriate status code (e.g., 207 Multi-Status if partial success)
	statusCode := http.StatusOK // Or http.StatusNoContent if all succeed and nothing to return
	if len(resp.FailedIDs) > 0 {
		statusCode = http.StatusMultiStatus
	}
	c.JSON(statusCode, resp)
}

// --- Aggregation & Pivot Handlers ---

// AggregateRecords 处理聚合查询请求
func (h *RecordHandler) AggregateRecords(c *gin.Context) {
	var req model.AggregationRequest
	// We might bind from query parameters or body depending on API design
	if err := c.ShouldBindJSON(&req); err != nil { // Assuming JSON body for now
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	// Ensure table_id is present (might come from path or body)
	if req.TableID == "" {
		pathTableID := c.Param("table_id")
		if pathTableID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "table_id is required in path or body"})
			return
		}
		req.TableID = pathTableID
	}

	resp, err := h.RecordService.AggregateRecords(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to aggregate records: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// PivotRecords 处理透视表查询请求
func (h *RecordHandler) PivotRecords(c *gin.Context) {
	var req model.PivotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	// Ensure table_id is present
	if req.TableID == "" {
		pathTableID := c.Param("table_id")
		if pathTableID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "table_id is required in path or body"})
			return
		}
		req.TableID = pathTableID
	}

	resp, err := h.RecordService.PivotRecords(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate pivot table: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
