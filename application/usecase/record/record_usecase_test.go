package record_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"theta-be/application/usecase/record"
	"theta-be/domain/entity"
	"theta-be/domain/enum"
	"theta-be/pkg/errs"
)

type mockRecordRepository struct {
	records map[string]*entity.Record
}

func newMockRepository() *mockRecordRepository {
	return &mockRecordRepository{
		records: make(map[string]*entity.Record),
	}
}

func (m *mockRecordRepository) Create(ctx context.Context, r *entity.Record) error {
	m.records[r.ID] = r
	return nil
}

func (m *mockRecordRepository) FindByID(ctx context.Context, id string) (*entity.Record, error) {
	r, exists := m.records[id]
	if !exists {
		return nil, errs.ErrNotFound
	}
	return r, nil
}

func (m *mockRecordRepository) FindAll(ctx context.Context, offset, limit int) ([]*entity.Record, int64, error) {
	var list []*entity.Record
	for _, r := range m.records {
		list = append(list, r)
	}
	return list, int64(len(list)), nil
}

func (m *mockRecordRepository) Update(ctx context.Context, r *entity.Record) error {
	if _, exists := m.records[r.ID]; !exists {
		return errs.ErrNotFound
	}
	m.records[r.ID] = r
	return nil
}

func (m *mockRecordRepository) Delete(ctx context.Context, id string) error {
	if _, exists := m.records[id]; !exists {
		return errs.ErrNotFound
	}
	delete(m.records, id)
	return nil
}

func TestRecordUsecase_CreateRecord(t *testing.T) {
	fixedTime := time.Date(2026, 10, 14, 12, 0, 0, 0, time.UTC)
	fixedID := "test-record-uuid-1"

	t.Run("SC-01: Valid input creates record in DRAFT status", func(t *testing.T) {
		repo := newMockRepository()
		uc := record.NewUsecaseWithClock(repo, func() string { return fixedID }, func() time.Time { return fixedTime })

		res, err := uc.CreateRecord(context.Background(), record.CreateRecordInput{
			Title:       "Clean Architecture Layering",
			Description: "SALT Boilerplate",
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, fixedID, res.ID)
		assert.Equal(t, "Clean Architecture Layering", res.Title)
		assert.Equal(t, enum.StatusDraft, res.Status)
		assert.Equal(t, fixedTime, res.CreatedAt)
	})

	t.Run("SC-02: Empty title returns ErrInvalidInput", func(t *testing.T) {
		repo := newMockRepository()
		uc := record.NewUsecase(repo)

		res, err := uc.CreateRecord(context.Background(), record.CreateRecordInput{
			Title: "   ",
		})

		assert.ErrorIs(t, err, errs.ErrInvalidInput)
		assert.Nil(t, res)
	})
}

func TestRecordUsecase_GetRecord(t *testing.T) {
	repo := newMockRepository()
	uc := record.NewUsecase(repo)

	created, _ := uc.CreateRecord(context.Background(), record.CreateRecordInput{
		Title: "Existing Item",
	})

	t.Run("SC-03: Existing ID returns record", func(t *testing.T) {
		res, err := uc.GetRecord(context.Background(), created.ID)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, created.ID, res.ID)
	})

	t.Run("SC-04: Non-existent ID returns ErrNotFound", func(t *testing.T) {
		res, err := uc.GetRecord(context.Background(), "unknown-id")
		assert.ErrorIs(t, err, errs.ErrNotFound)
		assert.Nil(t, res)
	})
}

func TestRecordUsecase_ActivateRecord(t *testing.T) {
	t.Run("SC-05: Valid transition from DRAFT to ACTIVE", func(t *testing.T) {
		repo := newMockRepository()
		uc := record.NewUsecase(repo)

		created, err := uc.CreateRecord(context.Background(), record.CreateRecordInput{
			Title: "Draft Item",
		})
		assert.NoError(t, err)

		activated, err := uc.ActivateRecord(context.Background(), created.ID)
		assert.NoError(t, err)
		assert.Equal(t, enum.StatusActive, activated.Status)
	})

	t.Run("SC-06: Transition from ARCHIVED to ACTIVE fails", func(t *testing.T) {
		repo := newMockRepository()
		uc := record.NewUsecase(repo)

		created, _ := uc.CreateRecord(context.Background(), record.CreateRecordInput{
			Title: "Archived Item",
		})
		_ = created.Archive(time.Now())
		_ = repo.Update(context.Background(), created)

		activated, err := uc.ActivateRecord(context.Background(), created.ID)
		assert.ErrorIs(t, err, errs.ErrInvalidStateTransition)
		assert.Nil(t, activated)
	})
}

func TestRecordUsecase_DeleteRecord(t *testing.T) {
	repo := newMockRepository()
	uc := record.NewUsecase(repo)

	created, _ := uc.CreateRecord(context.Background(), record.CreateRecordInput{
		Title: "Item To Delete",
	})

	t.Run("SC-07: Delete existing record succeeds", func(t *testing.T) {
		err := uc.DeleteRecord(context.Background(), created.ID)
		assert.NoError(t, err)

		_, err = uc.GetRecord(context.Background(), created.ID)
		assert.ErrorIs(t, err, errs.ErrNotFound)
	})

	t.Run("SC-08: Delete non-existent record returns ErrNotFound", func(t *testing.T) {
		err := uc.DeleteRecord(context.Background(), "non-existent-id")
		assert.ErrorIs(t, err, errs.ErrNotFound)
	})
}
