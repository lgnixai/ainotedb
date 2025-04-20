package router

import (
	"github.com/gin-gonic/gin"

	"github.com/undb/undb-go/internal/record/handler"
)

// RegisterRecordRoutes registers record-related routes.  Improved structure from edited code, but retains original functionality.

func RegisterRoutes(r *gin.RouterGroup, recordHandler *handler.RecordHandler) {
	recordGroup := r.Group("/records")
	{
		records := recordGroup.Group("") // Use nested group for better organization
		{
			records.POST("", recordHandler.CreateRecord)
			records.GET("/:id", recordHandler.GetRecord)
			records.PUT("/:id", recordHandler.UpdateRecord)
			records.DELETE("/:id", recordHandler.DeleteRecord)

			records.POST("/batch-create", recordHandler.BatchCreateRecords)
			records.POST("/batch-update", recordHandler.BatchUpdateRecords)
			records.POST("/batch-delete", recordHandler.BatchDeleteRecords)

			records.POST("/aggregate", recordHandler.AggregateRecords)
			records.POST("/pivot", recordHandler.PivotRecords)

			// Retain functionality for getting records by table ID using query parameter - requires modification to handler function.
			// This is a compromise as the edited code lacks this functionality.
			// Could consider a different route structure if the number of table-specific operations grows.
			records.GET("", recordHandler.GetRecordsByTable)

		}
	}
	// Separate route for getting records by table ID using path parameter, as in the original code
	//router.GET("/api/tables/:table_id/records", recordHandler.GetRecordsByTable)
}
