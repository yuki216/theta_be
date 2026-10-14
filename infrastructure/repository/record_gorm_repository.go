package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"theta-be/domain/entity"
	"theta-be/domain/enum"
	domainrepo "theta-be/domain/repository"
	"theta-be/pkg/errs"
)

type RecordGormModel struct {
	ID          string     `gorm:"primaryKey;type:varchar(36)"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:text"`
	Status      string     `gorm:"type:varchar(32);not null;default:'DRAFT'"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `gorm:"index"`
}

func (RecordGormModel) TableName() string {
	return "theta_records"
}

func toDomain(m *RecordGormModel) *entity.Record {
	return &entity.Record{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Status:      enum.Status(m.Status),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt,
	}
}

func fromDomain(e *entity.Record) *RecordGormModel {
	return &RecordGormModel{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		Status:      e.Status.String(),
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   e.DeletedAt,
	}
}

type recordGormRepository struct {
	db *gorm.DB
}

func NewRecordGormRepository(db *gorm.DB) domainrepo.RecordRepository {
	_ = db.AutoMigrate(&RecordGormModel{})
	return &recordGormRepository{db: db}
}

func (r *recordGormRepository) Create(ctx context.Context, record *entity.Record) error {
	model := fromDomain(record)
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *recordGormRepository) FindByID(ctx context.Context, id string) (*entity.Record, error) {
	var model RecordGormModel
	err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}
	return toDomain(&model), nil
}

func (r *recordGormRepository) FindAll(ctx context.Context, offset, limit int) ([]*entity.Record, int64, error) {
	var models []*RecordGormModel
	var total int64

	if err := r.db.WithContext(ctx).Model(&RecordGormModel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&models).Error
	if err != nil {
		return nil, 0, err
	}

	records := make([]*entity.Record, 0, len(models))
	for _, m := range models {
		records = append(records, toDomain(m))
	}

	return records, total, nil
}

func (r *recordGormRepository) Update(ctx context.Context, record *entity.Record) error {
	model := fromDomain(record)
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *recordGormRepository) Delete(ctx context.Context, id string) error {
	res := r.db.WithContext(ctx).Delete(&RecordGormModel{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errs.ErrNotFound
	}
	return nil
}
