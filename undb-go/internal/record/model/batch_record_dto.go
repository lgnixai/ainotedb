package model

// BatchUpdateRecordRequest 定义了批量更新记录的请求结构

// UpdateRecordData 定义了单条记录更新的数据结构
type UpdateRecordData struct {
	ID   string                 `json:"id" binding:"required"`   // 记录ID
	Data map[string]interface{} `json:"data" binding:"required"` // 需要更新的数据
}

// BatchUpdateRecordResponse 定义了批量更新记录的响应结构
