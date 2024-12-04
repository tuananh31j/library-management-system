package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/errorCustom"
)

func Error(c *fiber.Ctx, statusCode int, message string) error {
	errRes := c.Status(statusCode).JSON(errorCustom.ErrorResponse{
		Code:    statusCode,
		Status:  "error",
		Message: message,
	})

	return errRes
}
