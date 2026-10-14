package repository

import (
	"context"
	"sort"
	"sync"

	"theta-be/domain/entity"
	domainrepo "theta-be/domain/repository"
	"theta-be/pkg/errs"
)

type recordMemoryRepository struct {
	mu      sync.RWMutex
	records map[string]*entity.Record
}

func NewRecordMemoryRepository() domainrepo.RecordRepository {
	return &recordMemoryRepository{
		records: make(map[string]*entity.Record),
	}
}

func (r *recordMemoryRepository) Create(ctx context.Context, record *entity.Record) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.records[record.ID]; exists {
		return errs.ErrConflict
	}
	clone := *record
	r.records[record.ID] = &clone
	return nil
}

func (r *recordMemoryRepository) FindByID(ctx context.Context, id string) (*entity.Record, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rec, exists := r.records[id]
	if !exists {
		return nil, errs.ErrNotFound
	}
	clone := *rec
	return &clone, nil
}

func (r *recordMemoryRepository) FindAll(ctx context.Context, offset, limit int) ([]*entity.Record, int64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var list []*entity.Record
	for _, rec := range r.records {
		clone := *rec
		list = append(list, &clone)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.After(list[j].CreatedAt)
	})

	total := int64(len(list))
	if offset >= len(list) {
		return []*entity.Record{}, total, nil
	}

	end := offset + limit
	if end > len(list) {
		end = len(list)
	}

	return list[offset:end], total, nil
}

func (r *recordMemoryRepository) Update(ctx context.Context, record *entity.Record) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.records[record.ID]; !exists {
		return errs.ErrNotFound
	}
	clone := *record
	r.records[record.ID] = &clone
	return nil
}

func (r *recordMemoryRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.records[id]; !exists {
		return errs.ErrNotFound
	}
	delete(r.records, id)
	return nil
}
