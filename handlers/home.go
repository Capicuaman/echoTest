// handlers/home.go
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// func HomeHandler(c echo.Context) error {
// 	return c.Render(http.StatusOK, "home.html", nil)
// }

func HomeHandler(c echo.Context) error {
	data := map[string]interface{}{
		"Title":       "Home Page",
		"PageTitle":   "Welcome to our Home Page",
		"PageContent": "This is the home page content.",
	}
	return c.Render(http.StatusOK, "home.html", data)
}
