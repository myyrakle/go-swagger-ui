package fiberv3swagger

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	swagger "github.com/myyrakle/go-swagger-ui"
)

// Serve registers Swagger UI handlers on the Fiber app under prefix.
//
// It serves the embedded Swagger UI index page, static assets, and the spec
// document from prefix. specBytes is used to choose the index page and doc
// endpoint: valid JSON specs load ./doc.json, and all other specs load
// ./doc.yaml.
func Serve(app *fiber.App, prefix string, specBytes []byte) {
	isJSON := json.Valid(specBytes)
	prefix = normalizePrefix(prefix)

	if prefix != "/" {
		base := strings.TrimSuffix(prefix, "/")
		app.Get(base, func(c fiber.Ctx) error {
			if c.Path() != base {
				return c.Next()
			}
			return c.Redirect().Status(http.StatusMovedPermanently).To(prefix)
		})
	}

	app.Get(prefix+"*", func(c fiber.Ctx) error {
		file := strings.TrimPrefix(c.Params("*"), "/")

		switch file {
		case "":
			fallthrough
		case "index.html":
			if isJSON {
				return c.Type("html", "utf-8").Send([]byte(swagger.JSONIndexHTML))
			}
			return c.Type("html", "utf-8").Send([]byte(swagger.YAMLIndexHTML))
		case "doc.json":
			if !isJSON {
				return c.SendStatus(http.StatusNotFound)
			}
			docBytes := specBytes
			if requestScheme(c) == swagger.SchemeHTTP {
				docBytes, _ = swagger.PreprocessSchemeInJSONSpec(specBytes, swagger.SchemeHTTP)
			} else {
				docBytes, _ = swagger.PreprocessSchemeInJSONSpec(specBytes, swagger.SchemeHTTPS)
			}
			return c.Type("json", "utf-8").Send(docBytes)
		case "doc.yaml":
			if isJSON {
				return c.SendStatus(http.StatusNotFound)
			}
			docBytes := specBytes
			if requestScheme(c) == swagger.SchemeHTTP {
				docBytes, _ = swagger.PreprocessSchemeInYAMLSpec(specBytes, swagger.SchemeHTTP)
			} else {
				docBytes, _ = swagger.PreprocessSchemeInYAMLSpec(specBytes, swagger.SchemeHTTPS)
			}
			c.Set(fiber.HeaderContentType, "application/x-yaml; charset=utf-8")
			return c.Send(docBytes)
		case "swagger-ui.css":
			return c.Type("css", "utf-8").Send([]byte(swagger.SwaggerUICSS))
		case "index.css":
			return c.Type("css", "utf-8").Send([]byte(swagger.IndexCSS))
		case "swagger-ui-bundle.js":
			return c.Type("js", "utf-8").Send([]byte(swagger.SwaggerUIBundleJS))
		case "swagger-ui-standalone-preset.js":
			return c.Type("js", "utf-8").Send([]byte(swagger.SwaggerUIStandalonePresetJS))
		case "favicon-16x16.png":
			return c.Type("png").Send(swagger.Favicon16PNG)
		case "favicon-32x32.png":
			return c.Type("png").Send(swagger.Favicon32PNG)
		default:
			return c.SendStatus(http.StatusNotFound)
		}
	})
}

func requestScheme(c fiber.Ctx) swagger.Scheme {
	switch c.Get("X-Forwarded-Proto") {
	case "http":
		return swagger.SchemeHTTP
	case "https":
		return swagger.SchemeHTTPS
	}
	if c.Secure() {
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
