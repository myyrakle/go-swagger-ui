package swagger

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	swagger "github.com/myyrakle/go-swagger-ui"
)

type SpecFormat string

const (
	JSON SpecFormat = "json"
	YAML SpecFormat = "yaml"
)

type SpecSource interface {
	Format() SpecFormat
	Content() []byte
}

type SpecSourceJSON struct {
	JSON []byte
}

func (s SpecSourceJSON) Format() SpecFormat {
	return JSON
}

func (s SpecSourceJSON) Content() []byte {
	return s.JSON
}

type SpecSourceYAML struct {
	YAML []byte
}

func (s SpecSourceYAML) Format() SpecFormat {
	return YAML
}

func (s SpecSourceYAML) Content() []byte {
	return s.YAML
}

func Setup(e *echo.Echo, prefix string, format SpecFormat) {
	e.GET(prefix+"*",
		func(c echo.Context) error {
			file := c.Param("*")
			if file == "" {
				base := strings.TrimSuffix(c.Request().URL.Path, "/")
				return c.Redirect(http.StatusMovedPermanently, base+"/index.html")
			}

			switch file {
			case "index.html":
				switch format {
				case JSON:
					return c.HTMLBlob(http.StatusOK, []byte(swagger.JSONIndexHTML))
				case YAML:
					return c.HTMLBlob(http.StatusOK, []byte(swagger.YAMLIndexHTML))
				}
			case "swagger-ui.css":
				return c.Blob(http.StatusOK, "text/css; charset=utf-8", []byte(swagger.SwaggerUICSS))
			case "index.css":
				return c.Blob(http.StatusOK, "text/css; charset=utf-8", []byte(swagger.IndexCSS))
			case "swagger-ui-bundle.js":
				return c.Blob(http.StatusOK, "application/javascript; charset=utf-8", []byte(swagger.SwaggerUIBundleJS))
			case "swagger-ui-standalone-preset.js":
				return c.Blob(http.StatusOK, "application/javascript; charset=utf-8", []byte(swagger.SwaggerUIStandalonePresetJS))
			case "favicon-16x16.png":
				return c.Blob(http.StatusOK, "image/png", swagger.Favicon16PNG)
			case "favicon-32x32.png":
				return c.Blob(http.StatusOK, "image/png", swagger.Favicon32PNG)
			}

			return echo.ErrNotFound
		},
	)
}
