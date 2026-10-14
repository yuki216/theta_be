package httpresponse

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"theta-be/pkg/errs"
)

type Envelope struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Error   *ErrorBody  `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ErrorBody struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func Success(c *gin.Context, statusCode int, data interface{}, meta interface{}) {
	c.JSON(statusCode, Envelope{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func SuccessMessage(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Envelope{
		Success: true,
		Message: message,
	})
}

func Error(c *gin.Context, statusCode int, code string, message string, details interface{}) {
	c.JSON(statusCode, Envelope{
		Success: false,
		Error: &ErrorBody{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

func HandleDomainError(c *gin.Context, err error) {
	var codeErr *errs.CodeError
	if errors.As(err, &codeErr) {
		status := http.StatusBadRequest
		switch codeErr.Code {
		case "RESOURCE_NOT_FOUND":
			status = http.StatusNotFound
		case "RESOURCE_ALREADY_EXISTS":
			status = http.StatusConflict
		case "INVALID_STATE_TRANSITION":
			status = http.StatusUnprocessableEntity
		case "UNAUTHORIZED":
			status = http.StatusUnauthorized
		case "FORBIDDEN":
			status = http.StatusForbidden
		}
		Error(c, status, codeErr.Code, codeErr.Message, nil)
		return
	}

	switch {
	case errors.Is(err, errs.ErrInvalidInput):
		Error(c, http.StatusBadRequest, "INVALID_INPUT", err.Error(), nil)
	case errors.Is(err, errs.ErrNotFound):
		Error(c, http.StatusNotFound, "RESOURCE_NOT_FOUND", err.Error(), nil)
	case errors.Is(err, errs.ErrConflict):
		Error(c, http.StatusConflict, "RESOURCE_ALREADY_EXISTS", err.Error(), nil)
	case errors.Is(err, errs.ErrInvalidStateTransition):
		Error(c, http.StatusUnprocessableEntity, "INVALID_STATE_TRANSITION", err.Error(), nil)
	case errors.Is(err, errs.ErrUnauthorized):
		Error(c, http.StatusUnauthorized, "UNAUTHORIZED", err.Error(), nil)
	case errors.Is(err, errs.ErrForbidden):
		Error(c, http.StatusForbidden, "FORBIDDEN", err.Error(), nil)
	default:
		Error(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "An unexpected error occurred", nil)
	}
}
