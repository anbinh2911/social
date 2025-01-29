package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userGroup := e.Group("/users")
	userGroup.POST("/v1/regsiter", s.registerUserHandler)
	userGroup.POST("/v1/login", s.loginHandler)
	userGroup.POST("/v1/forgot", s.forgotPasswordHandler)
	userGroup.GET("/v1/detail", s.userDetailHandler)

	return e
}
