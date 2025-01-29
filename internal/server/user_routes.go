package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) registerUserHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Register new user",
	}

	return c.JSON(http.StatusOK, resp)

}

func (s *Server) loginHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Register new user",
	}

	return c.JSON(http.StatusOK, resp)

}

func (s *Server) forgotPasswordHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Register new user",
	}

	return c.JSON(http.StatusOK, resp)

}

func (s *Server) userDetailHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Register new user",
	}

	return c.JSON(http.StatusOK, resp)

}
