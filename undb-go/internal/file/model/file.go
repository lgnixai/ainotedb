
package model

import (
	"time"
	"mime/multipart"
)

type File struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	MimeType  string    `json:"mime_type"`
	Path      string    `json:"path"`
	RecordID  string    `json:"record_id"`
	TableID   string    `json:"table_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FileUploadRequest struct {
	File      *multipart.FileHeader `form:"file" binding:"required"`
	RecordID  string               `form:"record_id" binding:"required"`
	TableID   string               `form:"table_id" binding:"required"`
}

type FileResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Size     int64  `json:"size"` 
	MimeType string `json:"mime_type"`
	URL      string `json:"url"`
}
