package swagger

import (
	"net/http"

	"github.com/labstack/echo/v4"
	swagger "github.com/myyrakle/go-swagger-ui"
)

func Serve(e *echo.Echo, prefix string, specBytes []byte) {
	contentType := http.DetectContentType(specBytes)

	e.GET(prefix+"*",
		func(c echo.Context) error {
			file := c.Param("*")

			switch file {
			case "":
				fallthrough
			case "index.html":

				switch contentType {
				case "application/json":
					return c.HTMLBlob(http.StatusOK, []byte(swagger.JSONIndexHTML))
				default:
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
