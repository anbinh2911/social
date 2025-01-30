package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func RequestIdMiddleware() fiber.Handler {
	return requestid.New(requestid.Config{
		Header: "X-Request-Id",
		Generator: func() string {
			return "static-id"
		},
	})
}
