package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Routing struct {
	//! YOUR ROUTING STRUCT
}

func RegisterRouters(e *echo.Echo, routing *Routing) {
	e.Use(middleware.CORS())

	homeGroup := e.Group("/v1/")
	homeGroup.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "HELLO WORLD")
	})

	//! YOUR ROUTING GROUPS
}
