
package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"github.com/google/uuid"
	"github.com/undb/undb-go/internal/file/model"
)

type FileService interface {
	Upload(ctx context.Context, req *model.FileUploadRequest) (*model.FileResponse, error)
	Download(ctx context.Context, id string) (*os.File, error)
	Delete(ctx context.Context, id string) error
}

type fileService struct {
	uploadDir string
	repo     FileRepository
}

func NewFileService(uploadDir string, repo FileRepository) FileService {
	return &fileService{
		uploadDir: uploadDir,
		repo:     repo,
	}
}

func (s *fileService) Upload(ctx context.Context, req *model.FileUploadRequest) (*model.FileResponse, error) {
	file := req.File
	
	// Generate unique ID and file path
	id := uuid.New().String()
	ext := filepath.Ext(file.Filename)
	path := filepath.Join(s.uploadDir, id+ext)

	// Create file record
	fileModel := &model.File{
		ID:       id,
		Name:     file.Filename,
		Size:     file.Size,
		MimeType: file.Header.Get("Content-Type"),
		Path:     path,
		RecordID: req.RecordID,
		TableID:  req.TableID,
	}

	// Save file to disk
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, fmt.Errorf("failed to copy file: %w", err)
	}

	// Save to database
	if err := s.repo.Create(ctx, fileModel); err != nil {
		// Cleanup file if database insert fails
		os.Remove(path)
		return nil, fmt.Errorf("failed to save file record: %w", err)
	}

	return &model.FileResponse{
		ID:       fileModel.ID,
		Name:     fileModel.Name,
		Size:     fileModel.Size,
		MimeType: fileModel.MimeType,
		URL:      fmt.Sprintf("/api/files/%s", fileModel.ID),
	}, nil
}

func (s *fileService) Download(ctx context.Context, id string) (*os.File, error) {
	file, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("file not found: %w", err)
	}

	return os.Open(file.Path)
}

func (s *fileService) Delete(ctx context.Context, id string) error {
	file, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("file not found: %w", err)
	}

	if err := os.Remove(file.Path); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return s.repo.Delete(ctx, id)
}
