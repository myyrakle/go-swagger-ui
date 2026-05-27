package ginv1swagger

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	swagger "github.com/myyrakle/go-swagger-ui"
)

// Serve registers Swagger UI handlers on the Gin engine under prefix.
//
// It serves the embedded Swagger UI index page, static assets, and the spec
// document from prefix. specBytes is used to choose the index page and doc
// endpoint: valid JSON specs load ./doc.json, and all other specs load
// ./doc.yaml.
func Serve(e *gin.Engine, prefix string, specBytes []byte) {
	isJSON := json.Valid(specBytes)
	prefix = normalizePrefix(prefix)

	if prefix != "/" {
		e.GET(strings.TrimSuffix(prefix, "/"), func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, prefix)
		})
	}

	e.GET(prefix+"*file", func(c *gin.Context) {
		file := strings.TrimPrefix(c.Param("file"), "/")

		switch file {
		case "":
			fallthrough
		case "index.html":
			if isJSON {
				c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(swagger.JSONIndexHTML))
				return
			}
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(swagger.YAMLIndexHTML))
		case "doc.json":
			if !isJSON {
				c.Status(http.StatusNotFound)
				return
			}
			docBytes := specBytes
			if requestScheme(c) == swagger.SchemeHTTP {
				docBytes, _ = swagger.PreprocessSchemeInJSONSpec(specBytes, swagger.SchemeHTTP)
			} else {
				docBytes, _ = swagger.PreprocessSchemeInJSONSpec(specBytes, swagger.SchemeHTTPS)
			}
			c.Data(http.StatusOK, "application/json; charset=utf-8", docBytes)
		case "doc.yaml":
			if isJSON {
				c.Status(http.StatusNotFound)
				return
			}
			docBytes := specBytes
			if requestScheme(c) == swagger.SchemeHTTP {
				docBytes, _ = swagger.PreprocessSchemeInYAMLSpec(specBytes, swagger.SchemeHTTP)
			} else {
				docBytes, _ = swagger.PreprocessSchemeInYAMLSpec(specBytes, swagger.SchemeHTTPS)
			}
			c.Data(http.StatusOK, "application/x-yaml; charset=utf-8", docBytes)
		case "swagger-ui.css":
			c.Data(http.StatusOK, "text/css; charset=utf-8", []byte(swagger.SwaggerUICSS))
		case "index.css":
			c.Data(http.StatusOK, "text/css; charset=utf-8", []byte(swagger.IndexCSS))
		case "swagger-ui-bundle.js":
			c.Data(http.StatusOK, "application/javascript; charset=utf-8", []byte(swagger.SwaggerUIBundleJS))
		case "swagger-ui-standalone-preset.js":
			c.Data(http.StatusOK, "application/javascript; charset=utf-8", []byte(swagger.SwaggerUIStandalonePresetJS))
		case "favicon-16x16.png":
			c.Data(http.StatusOK, "image/png", swagger.Favicon16PNG)
		case "favicon-32x32.png":
			c.Data(http.StatusOK, "image/png", swagger.Favicon32PNG)
		default:
			c.Status(http.StatusNotFound)
		}
	})
}

func requestScheme(c *gin.Context) swagger.Scheme {
	switch c.Request.Header.Get("X-Forwarded-Proto") {
	case "http":
		return swagger.SchemeHTTP
	case "https":
		return swagger.SchemeHTTPS
	}
	if c.Request.TLS != nil {
		return swagger.SchemeHTTPS
	}
	return swagger.SchemeHTTP
}

func normalizePrefix(prefix string) string {
	if prefix == "" || prefix == "/" {
		return "/"
	}
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	return strings.TrimRight(prefix, "/") + "/"
}
