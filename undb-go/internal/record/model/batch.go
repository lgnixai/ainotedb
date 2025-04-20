package model

type BatchRecordRequest struct {
    Records []*Record `json:"records" binding:"required,dive"`
}

type BatchDeleteRequest struct {
    IDs []string `json:"ids" binding:"required,min=1"`
}
