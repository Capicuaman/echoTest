// handlers/about.go
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// func AboutHandler(c echo.Context) error {
// 	return c.Render(http.StatusOK, "about.html", nil)
// }

func AboutHandler(c echo.Context) error {
	data := map[string]interface{}{
		"Title":       "about",
		"PageTitle":   "Welcome to our About Page",
		"PageContent": "This is the ABOUT page content.",
	}
	return c.Render(http.StatusOK, "about.html", data)
}
