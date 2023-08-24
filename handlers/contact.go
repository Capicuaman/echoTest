// handlers/contact.go
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ContactHandler(c echo.Context) error {
	data := map[string]interface{}{
		"Title":       "Contact",
		"PageTitle":   "Welcome to our CONTACT page ",
		"PageContent": "This is CONTENT page content.",
	}
	return c.Render(http.StatusOK, "contact.html", data)
}
