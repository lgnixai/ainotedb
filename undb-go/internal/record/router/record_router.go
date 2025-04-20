package router

import (
	"github.com/gin-gonic/gin"

	"github.com/undb/undb-go/internal/record/handler"
)

// RegisterRecordRoutes registers record-related routes.  Improved structure from edited code, but retains original functionality.

import "github.com/undb/undb-go/internal/user/middleware"

func RegisterRoutes(r *gin.RouterGroup, recordHandler *handler.RecordHandler) {
	recordGroup := r.Group("/records")
	recordGroup.Use(middleware.AuthMiddleware())
	{
		records := recordGroup.Group("")
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
			records.GET("", recordHandler.GetRecordsByTable)
		}
	}
}
