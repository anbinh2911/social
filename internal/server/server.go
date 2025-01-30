package server

import (
	"github.com/anbinh2911/social/internal/database"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	*fiber.App

	db database.Service
}

func NewServer() *Server {
	server := &Server{
		App: fiber.New(fiber.Config{
			ServerHeader: "Social",
			AppName:      "Social",
		}),

		db: database.New(),
	}

	return server
}
