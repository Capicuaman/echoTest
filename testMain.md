// Package Declaration: Every Go file starts with a package declaration, indicating which package the file belongs to. Packages are a way to organize and reuse code.
package main

// Import Statements: Import statements are used to include external packages or libraries in your code. The imported packages provide functionality that you can use in your program.
import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Define a template renderer
type TemplateRenderer struct {
	templates *template.Template
}

// Function Declaration: Functions are declared using the func keyword, followed by the function name, parameter list (if any), return type (if any), and the function body.
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)

}

// Main Function: Every Go program must have a main function. It's the entry point of the program and is executed when the program starts.
func main() {
	// Create an echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//  Initialize the template Renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Define routes and handlers
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.Render(http.StatusOK, "hello.html", map[string]interface{}{
			"name": name,
		})
	})
	// Start the server
	e.Logger.Fatal(e.Start(":8080"))

	// GO COMMON STRUCTURES

	// Variables
	message := "CAPICUAMAN"
	fmt.Println(message)

	age := 25
	day := "monday"

	// Constants
	const pi = 3.14159

	// Control Structures: Go supports typical control structures like if, for, and switch.
	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	switch day {
	case "Monday":
		fmt.Println("It's the start of the week.")
	case "Friday":
		fmt.Println("TGIF!")
	default:
		fmt.Println("It's another day.")
	}
	// Arrays and Slices: Arrays have a fixed size, while slices are dynamically resizable. Slices are more commonly used in Go.
	numbers := []int{1, 2, 3, 4, 5}
	// Maps: Maps are used to store key-value pairs.
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	// Structs: Structs allow you to define custom data types by grouping together different fields.

	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 26}
	fmt.Println(numbers, ages, p)

	// Methods: Go supports defining methods on types. Methods are functions that operate on specific types.

	// type Circle struct {
	// 	Radius float64
	// }
	// func (c Circle) Area() float64 {
	// 	return math.Pi * c.Radius * c.Radius
	// }

}
