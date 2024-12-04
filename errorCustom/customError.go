package errorCustom

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func BadRequest(msg string) ErrorResponse {
	if msg == "" {
		msg = "Something wrong!"
	}
	return ErrorResponse{
		Code:    fiber.StatusBadRequest,
		Status:  "Bad request",
		Message: msg,
	}
}
