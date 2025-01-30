package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/anbinh2911/social/internal/middlewares"
	"github.com/anbinh2911/social/internal/server"
)

func main() {

	server := server.NewServer()

	server.Use(middlewares.RequestIdMiddleware())
	server.Use(middlewares.LoggerMiddleware())
	server.Use(middlewares.RecoverMiddleware())

	server.RegisterRoutes()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
