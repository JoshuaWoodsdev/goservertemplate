import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"goservertemplate/pages"
)

// TemplateRenderer is a custom renderer to handle HTML templates
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders HTML templates
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Register HTML renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Routes
	e.GET("/", func(c echo.Context) error {
		data := pages.IndexPageData{
			Title:   "Welcome to My Echo App",
			Content: "This is the home page.",
		}
		return c.Render(http.StatusOK, "index.html", data)
	})

	e.GET("/pageone", func(c echo.Context) error {
		data := pages.PageOneData{
			Title:   "Page One",
			Content: "This is Page One.",
		}
		return c.Render(http.StatusOK, "pageone.html", data)
	})

	// Start the server on port 8080
	e.Start(":8080")
}