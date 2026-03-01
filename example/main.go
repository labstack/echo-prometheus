package main

import (
	"log/slog"
	"net/http"

	echoprometheus "github.com/labstack/echo-prometheus"
	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	// Enable metrics middleware
	e.Use(echoprometheus.NewMiddleware("myapp"))
	e.GET("/metrics", echoprometheus.NewHandler()) // expose metrics for scraping

	e.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	if err := e.Start(":8080"); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
