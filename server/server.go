package server

import (
	"fmt"
	"github.com/TianaNanta/web-echo/server/handlers"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	gowebly "github.com/gowebly/helpers"
)

// RunServer runs a new HTTP server with the loaded environment variables.
func RunServer() error {
	// Validate environment variables.
	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "7000"))
	if err != nil {
		return err
	}

	// Create a new Echo server.
	e := echo.New()

	// Add Echo middlewares.
	e.Use(middleware.Logger())

	// Handle static files.
	e.Static("/static", "./static")

	// Routes
	e.GET("/", handlers.IndexViewHandler)
	e.GET("/api/hello-world", handlers.ShowContentAPIHandler)

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      e, // handle all Echo routes
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Send log message.
	slog.Info("Starting server...", "port", port)

	return server.ListenAndServe()
}
