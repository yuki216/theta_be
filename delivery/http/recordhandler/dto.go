package recordhandler

import (
	"time"

	"theta-be/domain/entity"
	"theta-be/domain/enum"
)

type CreateRecordRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type RecordResponseDTO struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      enum.Status `json:"status"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type PaginationMeta struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}

func ToResponseDTO(r *entity.Record) RecordResponseDTO {
	return RecordResponseDTO{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		Status:      r.Status,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

func ToResponseDTOList(records []*entity.Record) []RecordResponseDTO {
	dtos := make([]RecordResponseDTO, 0, len(records))
	for _, r := range records {
		dtos = append(dtos, ToResponseDTO(r))
	}
	return dtos
}
