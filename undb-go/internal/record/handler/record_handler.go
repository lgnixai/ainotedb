package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/record/model"
	"github.com/undb/undb-go/internal/record/service"
)

type RecordHandler struct {
	recordService service.RecordService
}

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

	count, err := h.recordService.BatchCreate(c.Request.Context(), req.Records)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Records created successfully",
		"count":   count,
	})
}

func (h *RecordHandler) BatchUpdate(c *gin.Context) {
	var req model.BatchRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := h.recordService.BatchUpdate(c.Request.Context(), req.Records)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Records updated successfully",
		"count":   count,
	})
}

func (h *RecordHandler) BatchDelete(c *gin.Context) {
	var req model.BatchDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := h.recordService.BatchDelete(c.Request.Context(), req.IDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Records deleted successfully",
		"count":   count,
	})
}

func (h *RecordHandler) Aggregate(c *gin.Context) {
	var req model.AggregationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.recordService.Aggregate(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func (h *RecordHandler) Pivot(c *gin.Context) {
	var req model.PivotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.recordService.Pivot(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}