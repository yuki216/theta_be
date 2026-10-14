package entity

import (
	"strings"
	"time"

	"theta-be/domain/enum"
	"theta-be/pkg/errs"
)

type Record struct {
	ID          string
	Title       string
	Description string
	Status      enum.Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func NewRecord(id, title, description string, now time.Time) (*Record, error) {
	cleanTitle := strings.TrimSpace(title)
	if cleanTitle == "" {
		return nil, errs.ErrInvalidInput
	}
	return &Record{
		ID:          id,
		Title:       cleanTitle,
		Description: description,
		Status:      enum.StatusDraft,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (r *Record) Activate(now time.Time) error {
	if r.Status == enum.StatusArchived {
		return errs.ErrInvalidStateTransition
	}
	r.Status = enum.StatusActive
	r.UpdatedAt = now
	return nil
}

func (r *Record) Suspend(now time.Time) error {
	if r.Status != enum.StatusActive {
		return errs.ErrInvalidStateTransition
	}
	r.Status = enum.StatusSuspended
	r.UpdatedAt = now
	return nil
}

func (r *Record) Archive(now time.Time) error {
	if r.Status == enum.StatusArchived {
		return errs.ErrInvalidStateTransition
	}
	r.Status = enum.StatusArchived
	r.UpdatedAt = now
	return nil
}
