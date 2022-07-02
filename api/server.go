package main

import (
	"api/handler"
	"api/openapi"
	"api/store"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Store *store.Store
}

func startServer(c *Config, s *store.Store) error {
	e := echo.New()
	registerMiddlewares(e)

	handlers := handler.NewHandlers(s)

	openapi.RegisterHandlers(e, handlers)

	return e.Start(fmt.Sprintf(":%d", c.Server.Port))
}

func registerMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())
}
