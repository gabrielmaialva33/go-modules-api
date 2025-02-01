package exceptions

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// APIException represents a structured API error response
type APIException struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Error implements the error interface
func (e *APIException) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e.Status, e.Code, e.Message)
}

// NewAPIException creates a new APIException instance
func NewAPIException(status int, code, message string, details interface{}) *APIException {
	return &APIException{
		Status:  status,
		Code:    code,
		Message: message,
		Details: details,
	}
}

// Response returns a structured JSON response
func (e *APIException) Response(c *fiber.Ctx) error {
	return c.Status(e.Status).JSON(e)
}

func BadRequest(message string, details interface{}) *APIException {
	return NewAPIException(http.StatusBadRequest, "bad_request", message, details)
}

func NotFound(message string, details interface{}) *APIException {
	return NewAPIException(http.StatusNotFound, "not_found", message, details)
}

func InternalServerError(message string, details interface{}) *APIException {
	return NewAPIException(http.StatusInternalServerError, "internal_server_error", message, details)
}

func Conflict(message string, details interface{}) *APIException {
	return NewAPIException(http.StatusConflict, "duplicate_entry", message, details)
}
