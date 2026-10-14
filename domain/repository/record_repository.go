package repository

import (
	"context"

	"theta-be/domain/entity"
)

type RecordRepository interface {
	Create(ctx context.Context, record *entity.Record) error
	FindByID(ctx context.Context, id string) (*entity.Record, error)
	FindAll(ctx context.Context, offset, limit int) ([]*entity.Record, int64, error)
	Update(ctx context.Context, record *entity.Record) error
	Delete(ctx context.Context, id string) error
}
