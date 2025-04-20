package repository

import (
	"context"

	"github.com/undb/undb-go/internal/file/model"
)

type FileRepository interface {
	Create(ctx context.Context, file *model.File) error
	GetByID(ctx context.Context, id string) (*model.File, error)
	Delete(ctx context.Context, id string) error
}
