package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) registerUserHandler(c *fiber.Ctx) error {
	resp := map[string]string{
		"message": "Register new user",
	}

	return c.JSON(resp)
}

func (s *Server) loginHandler(c *fiber.Ctx) error {
	resp := map[string]string{
		"message": "Register new user",
	}

	return c.JSON(resp)
}

func (s *Server) forgotPasswordHandler(c *fiber.Ctx) error {
	resp := map[string]string{
		"message": "Register new user",
	}

	return c.JSON(resp)
}

func (s *Server) userDetailHandler(c *fiber.Ctx) error {
	resp := map[string]string{
		"message": "Register new user",
	}

	return c.JSON(resp)
}
