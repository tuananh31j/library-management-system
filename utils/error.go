package utils

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/errorCustom"
	"github.com/tuananh31j/library-management-system/response"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var httpError *fiber.Error
	if ok := errors.As(err, &httpError); ok {
		Log.Error(httpError.Message, httpError.Code)
		return response.Error(ctx, httpError.Code, httpError.Message)
	}

	// Kiểm tra nếu lỗi là lỗi tùy chỉnh
	var customError *errorCustom.ErrorResponse
	if ok := errors.As(err, &customError); ok {
		Log.Error(customError.Message, customError.Code)
		return response.Error(ctx, customError.Code, customError.Message)
	}
	log.Printf("Internal Error: %+v", err)

	return ctx.Status(fiber.StatusInternalServerError).JSON(errorCustom.ErrorResponse{
		Code:    fiber.StatusInternalServerError,
		Status:  "error",
		Message: "Internal server error",
	})
}
