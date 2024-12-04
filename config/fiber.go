package config

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/utils"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		DisableStartupMessage: false,
		ServerHeader:          "Fiber",
		AppName:               "library-management-system-api",
		ErrorHandler:          utils.ErrorHandler,
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
	}
}
