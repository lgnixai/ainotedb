package router

import (
	"github.com/gin-gonic/gin"
	"github.com/undb/undb-go/internal/record/handler"
)

// RegisterRecordRoutes 注册记录相关的路由
func RegisterRecordRoutes(router *gin.Engine, recordHandler *handler.RecordHandler) {
	// 定义记录相关的路由组
	// Consider prefixing with /api/v1 or similar
	recordGroup := router.Group("/api/records")
	{
		// Standard CRUD
		recordGroup.POST("", recordHandler.CreateRecord)
		recordGroup.GET("/:id", recordHandler.GetRecord) // Renamed from GetRecordByID for consistency
		// GET /table/:tableId seems redundant if GetRecordsByTable exists and uses path param
		// recordGroup.GET("/table/:tableId", recordHandler.GetRecordsByTableID)
		recordGroup.PUT("/:id", recordHandler.UpdateRecord)
		recordGroup.DELETE("/:id", recordHandler.DeleteRecord)

		// Get all records for a specific table
		// Use a more RESTful approach like /api/tables/:table_id/records
		// For now, keeping it under /api/records for simplicity, but needs table_id
		// This route might conflict with GET /:id if table_id is not handled carefully
		// Let's assume table_id is a query parameter for this route for now
		// recordGroup.GET("", recordHandler.GetRecordsByTable) // Needs modification if table_id is query param

		// Batch Operations
		recordGroup.POST("/batch-create", recordHandler.BatchCreateRecords)
		recordGroup.POST("/batch-update", recordHandler.BatchUpdateRecords) // Changed path from /batch
		recordGroup.POST("/batch-delete", recordHandler.BatchDeleteRecords)

		// Aggregation & Pivot (Could be POST or GET depending on complexity of request)
		// Using POST to allow for complex request bodies
		// These might be better under a specific table's route, e.g., /api/tables/:table_id/aggregate
		recordGroup.POST("/aggregate", recordHandler.AggregateRecords)
		recordGroup.POST("/pivot", recordHandler.PivotRecords)

		// Route for getting records by table ID (using path parameter)
		// Ensure this doesn't clash with other routes like /batch-create etc.
		// Maybe prefix table-specific routes? e.g., /table/:table_id/list
		router.GET("/api/tables/:table_id/records", recordHandler.GetRecordsByTable)

	}
}
