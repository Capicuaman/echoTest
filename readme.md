# Organizing a multi-page website using the Echo framework

structuring your routes, templates, and handlers effectively.

**1. Create a File Structure:**

Start by creating a well-organized file structure for your project. Here's an example structure:

```go
my-website/
  ├── main.go
  ├── handlers/
  │     ├── home.go
  │     ├── about.go
  │     └── contact.go
  ├── templates/
  │     ├── home.html
  │     ├── about.html
  │     ├── contact.html
  │     └── base.html
  └── static/
        ├── css/
        │     └── styles.css
        ├── images/
        └── js/
```

**2. Define Route Handlers:**

Create separate handler files in the `handlers` directory for each page of your website. For instance:

```go
// handlers/home.go
package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
    return c.Render(http.StatusOK, "home.html", nil)
}
```

```go
// handlers/about.go
package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func AboutHandler(c echo.Context) error {
    return c.Render(http.StatusOK, "about.html", nil)
}
```

```go
// handlers/contact.go
package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func ContactHandler(c echo.Context) error {
    return c.Render(http.StatusOK, "contact.html", nil)
}
```

**3. Define Templates:**

Create separate HTML template files for each page of your website in the `templates` directory. For example:

```html
<!-- templates/home.html -->
{{define "content"}}
<h1>Welcome to our Home Page</h1>
<p>This is the homepage of our website.</p>
{{end}}
```

```html
<!-- templates/about.html -->
{{define "content"}}
<h1>About Us</h1>
<p>Learn more about our company and mission.</p>
{{end}}
```

```html
<!-- templates/contact.html -->
{{define "content"}}
<h1>Contact Us</h1>
<p>Get in touch with us using the contact form below.</p>
{{end}}
```

**4. Define a Base Template:**

To achieve consistent styling across all pages, create a base template that includes common elements like the header, navigation, and footer. This template can be used as a layout for your other pages:

```html
<!-- templates/base.html -->
<!DOCTYPE html>
<html>
<head>
    <title>{{.title}}</title>
    <link rel="stylesheet" href="/static/css/styles.css">
</head>
<body>
    <header>
        <h1>My Website</h1>
        <nav>
            <a href="/">Home</a>
            <a href="/about">About</a>
            <a href="/contact">Contact</a>
        </nav>
    </header>
    <main>
        {{template "content" .}}
    </main>
    <footer>
        &copy; 2023 My Website
    </footer>
</body>
</html>
```

**5. Set Up Echo and Routes:**

In your `main.go` file, set up the Echo instance and define your routes using the handlers you've created:

```go
package main

import (
    "html/template"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "my-website/handlers"
)

type TemplateRenderer struct {
    templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    renderer := &TemplateRenderer{
        templates: template.Must(template.ParseGlob("templates/*.html")),
    }
    e.Renderer = renderer

    // Define routes
    e.GET("/", handlers.HomeHandler)
    e.GET("/about", handlers.AboutHandler)
    e.GET("/contact", handlers.ContactHandler)

    // Serve static files
    e.Static("/static", "static")

    e.Logger.Fatal(e.Start(":8080"))
}
```

**6. Run the Application:**

Navigate to your project directory and run the application:

```bash
go run main.go
```

**7. Access the Website:**

Open your web browser and access the following URLs:

- `http://localhost:8080` for the home page.
- `http://localhost:8080/about` for the about page.
- `http://localhost:8080/contact` for the contact page.

Each page will use the base template and render its specific content.

By organizing your project with separate handlers, templates, and static files, you create a maintainable and scalable structure for your multi-page website using the Echo framework.