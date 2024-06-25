package rendering

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer обеспечивает рендеринг HTML шаблонов
type TemplateRenderer struct {
	Templates *template.Template
}

// NewTemplateRenderer создает новый экземпляр TemplateRenderer
func NewTemplateRenderer(templateDir string) *TemplateRenderer {
	renderer := &TemplateRenderer{
		Templates: template.Must(template.ParseGlob(templateDir + "/*.html")),
	}
	return renderer
}

// Render реализует метод интерфейса echo.Renderer
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Пытаемся найти шаблон по имени
	tmpl := t.Templates.Lookup(name)
	if tmpl == nil {
		// Если не нашли шаблон, возвращаем ошибку
		return echo.NewHTTPError(http.StatusInternalServerError, "шаблон не найден")
	}

	// Если нашли, рендерим его с переданными данными
	return tmpl.ExecuteTemplate(w, "base.html", data)
}
