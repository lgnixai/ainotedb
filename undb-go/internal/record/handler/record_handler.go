package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/record/model"
	"github.com/undb/undb-go/internal/record/service"
)

type RecordHandler struct {
	recordService service.RecordService
}

// 路由兼容方法
func (h *RecordHandler) CreateRecord(c *gin.Context)    { h.Create(c) }
func (h *RecordHandler) GetRecord(c *gin.Context)      { h.GetByID(c) }
func (h *RecordHandler) UpdateRecord(c *gin.Context)   { h.Update(c) }
func (h *RecordHandler) DeleteRecord(c *gin.Context)   { h.Delete(c) }
func (h *RecordHandler) BatchCreateRecords(c *gin.Context)  { h.BatchCreate(c) }
func (h *RecordHandler) BatchUpdateRecords(c *gin.Context)  { h.BatchUpdate(c) }
func (h *RecordHandler) BatchDeleteRecords(c *gin.Context)  { h.BatchDelete(c) }
func (h *RecordHandler) AggregateRecords(c *gin.Context)    { h.Aggregate(c) }
func (h *RecordHandler) PivotRecords(c *gin.Context)        { h.Pivot(c) }
func (h *RecordHandler) GetRecordsByTable(c *gin.Context)   { h.GetByTableID(c) }

func NewRecordHandler(recordService service.RecordService) *RecordHandler {
	return &RecordHandler{recordService: recordService}
}

func (h *RecordHandler) Create(c *gin.Context) {
	var record model.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.recordService.Create(c.Request.Context(), &record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, record)
}

func (h *RecordHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	record, err := h.recordService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}

func (h *RecordHandler) GetByTableID(c *gin.Context) {
	tableID := c.Param("tableId")
	records, err := h.recordService.GetByTableID(c.Request.Context(), tableID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

func (h *RecordHandler) Update(c *gin.Context) {
	var record model.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.recordService.Update(c.Request.Context(), &record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, record)
}

func (h *RecordHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.recordService.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}

func (h *RecordHandler) BatchCreate(c *gin.Context) {
	var req model.BatchRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 需补充 TableID 字段，假设前端传递 table_id
// 将 []*model.Record 转换为 []map[string]interface{}
var recordDataList []map[string]interface{}
for _, r := range req.Records {
	recordDataList = append(recordDataList, r.Data)
}
tableIDStr := c.Query("table_id")
tableID, err := strconv.ParseUint(tableIDStr, 10, 64)
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid table_id"})
    return
}
resp, err := h.recordService.BatchCreateRecords(c.Request.Context(), model.BatchCreateRecordRequest{TableID: uint(tableID), Records: recordDataList})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *RecordHandler) BatchUpdate(c *gin.Context) {
	var req model.BatchRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 需转换为 []BatchUpdateRecordData
var updateRecords []model.BatchUpdateRecordData
for _, r := range req.Records {
	updateRecords = append(updateRecords, model.BatchUpdateRecordData{ID: strconv.FormatUint(uint64(r.ID), 10), Data: r.Data})
}
resp, err := h.recordService.BatchUpdateRecords(c.Request.Context(), model.BatchUpdateRecordRequest{Records: updateRecords})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *RecordHandler) BatchDelete(c *gin.Context) {
	var req model.BatchDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.recordService.BatchDeleteRecords(c.Request.Context(), model.BatchDeleteRecordRequest{RecordIDs: req.IDs})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *RecordHandler) Aggregate(c *gin.Context) {
	var req model.AggregationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.recordService.AggregateRecords(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *RecordHandler) Pivot(c *gin.Context) {
	var req model.PivotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.recordService.PivotRecords(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}