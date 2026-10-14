package recordhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"theta-be/application/usecase/record"
	"theta-be/pkg/httpresponse"
)

type Handler struct {
	usecase record.Usecase
}

func NewHandler(uc record.Usecase) *Handler {
	return &Handler{usecase: uc}
}

func (h *Handler) CreateRecord(c *gin.Context) {
	var req CreateRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresponse.Error(c, http.StatusBadRequest, "INVALID_INPUT", "Invalid JSON payload or missing title", err.Error())
		return
	}

	res, err := h.usecase.CreateRecord(c.Request.Context(), record.CreateRecordInput{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		httpresponse.HandleDomainError(c, err)
		return
	}

	httpresponse.Success(c, http.StatusCreated, ToResponseDTO(res), nil)
}

func (h *Handler) GetRecord(c *gin.Context) {
	id := c.Param("id")
	res, err := h.usecase.GetRecord(c.Request.Context(), id)
	if err != nil {
		httpresponse.HandleDomainError(c, err)
		return
	}

	httpresponse.Success(c, http.StatusOK, ToResponseDTO(res), nil)
}

func (h *Handler) ListRecords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	records, total, err := h.usecase.ListRecords(c.Request.Context(), page, limit)
	if err != nil {
		httpresponse.HandleDomainError(c, err)
		return
	}

	meta := PaginationMeta{
		Page:  page,
		Limit: limit,
		Total: total,
	}

	httpresponse.Success(c, http.StatusOK, ToResponseDTOList(records), meta)
}

func (h *Handler) ActivateRecord(c *gin.Context) {
	id := c.Param("id")
	res, err := h.usecase.ActivateRecord(c.Request.Context(), id)
	if err != nil {
		httpresponse.HandleDomainError(c, err)
		return
	}

	httpresponse.Success(c, http.StatusOK, ToResponseDTO(res), nil)
}

func (h *Handler) DeleteRecord(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteRecord(c.Request.Context(), id); err != nil {
		httpresponse.HandleDomainError(c, err)
		return
	}

	httpresponse.SuccessMessage(c, http.StatusOK, "Record deleted successfully")
}
