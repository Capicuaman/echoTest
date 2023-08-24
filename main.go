package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"echoTest/handlers" // Import your custom page handlers
)

// TemplateRenderer is a custom renderer for HTML templates
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders an HTML template
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New() // Create a new instance of Echo

	e.Use(middleware.Logger())  // Use middleware for logging
	e.Use(middleware.Recover()) // Use middleware for recovering from panics

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")), // Initialize the custom template renderer
	}
	e.Renderer = renderer // Set the custom template renderer for Echo

	// Define routes and handlers
	e.GET("/", handlers.HomeHandler)           // Handle the root URL with HomeHandler
	e.GET("/about", handlers.AboutHandler)     // Handle the "/about" URL with AboutHandler
	e.GET("/contact", handlers.ContactHandler) // Handle the "/contact" URL with ContactHandler

	// Serve static files (CSS, images, JS, etc.)
	e.Static("/static", "static")

	// Start the server
	e.Logger.Fatal(e.Start(":8080")) // Start the Echo server on port 8080
}
