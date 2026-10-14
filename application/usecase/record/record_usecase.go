package record

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"theta-be/domain/entity"
	"theta-be/domain/repository"
	"theta-be/pkg/errs"
)

type CreateRecordInput struct {
	Title       string
	Description string
}

type Usecase interface {
	CreateRecord(ctx context.Context, input CreateRecordInput) (*entity.Record, error)
	GetRecord(ctx context.Context, id string) (*entity.Record, error)
	ListRecords(ctx context.Context, page, limit int) ([]*entity.Record, int64, error)
	ActivateRecord(ctx context.Context, id string) (*entity.Record, error)
	DeleteRecord(ctx context.Context, id string) error
}

type usecase struct {
	repo    repository.RecordRepository
	idGen   func() string
	nowFunc func() time.Time
}

func NewUsecase(repo repository.RecordRepository) Usecase {
	return &usecase{
		repo: repo,
		idGen: func() string {
			return uuid.New().String()
		},
		nowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}
}

// NewUsecaseWithClock provides dependency injection for deterministic testing.
func NewUsecaseWithClock(repo repository.RecordRepository, idGen func() string, nowFunc func() time.Time) Usecase {
	return &usecase{
		repo:    repo,
		idGen:   idGen,
		nowFunc: nowFunc,
	}
}

func (u *usecase) CreateRecord(ctx context.Context, input CreateRecordInput) (*entity.Record, error) {
	title := strings.TrimSpace(input.Title)
	if title == "" {
		return nil, errs.ErrInvalidInput
	}

	rec, err := entity.NewRecord(u.idGen(), title, input.Description, u.nowFunc())
	if err != nil {
		return nil, err
	}

	if err := u.repo.Create(ctx, rec); err != nil {
		return nil, err
	}

	return rec, nil
}

func (u *usecase) GetRecord(ctx context.Context, id string) (*entity.Record, error) {
	cleanID := strings.TrimSpace(id)
	if cleanID == "" {
		return nil, errs.ErrInvalidInput
	}
	return u.repo.FindByID(ctx, cleanID)
}

func (u *usecase) ListRecords(ctx context.Context, page, limit int) ([]*entity.Record, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit
	return u.repo.FindAll(ctx, offset, limit)
}

func (u *usecase) ActivateRecord(ctx context.Context, id string) (*entity.Record, error) {
	rec, err := u.GetRecord(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := rec.Activate(u.nowFunc()); err != nil {
		return nil, err
	}

	if err := u.repo.Update(ctx, rec); err != nil {
		return nil, err
	}

	return rec, nil
}

func (u *usecase) DeleteRecord(ctx context.Context, id string) error {
	cleanID := strings.TrimSpace(id)
	if cleanID == "" {
		return errs.ErrInvalidInput
	}

	_, err := u.repo.FindByID(ctx, cleanID)
	if err != nil {
		return err
	}

	return u.repo.Delete(ctx, cleanID)
}
