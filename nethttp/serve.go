package nethttpswagger

import (
	"encoding/json"
	"net/http"
	"strings"

	swagger "github.com/myyrakle/go-swagger-ui"
)

// Handler returns an http.Handler that serves Swagger UI under prefix.
//
// It serves the embedded Swagger UI index page, static assets, and the spec
// document from prefix. specBytes is used to choose the index page and doc
// endpoint: valid JSON specs load ./doc.json, and all other specs load
// ./doc.yaml.
func Handler(prefix string, specBytes []byte) http.Handler {
	isJSON := json.Valid(specBytes)
	prefix = normalizePrefix(prefix)
	base := strings.TrimSuffix(prefix, "/")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if prefix != "/" && r.URL.Path == base {
			http.Redirect(w, r, prefix, http.StatusMovedPermanently)
			return
		}
		if !strings.HasPrefix(r.URL.Path, prefix) {
			http.NotFound(w, r)
			return
		}

		file := strings.TrimPrefix(r.URL.Path, prefix)
		switch file {
		case "":
			fallthrough
		case "index.html":
			if isJSON {
				write(w, http.StatusOK, "text/html; charset=utf-8", []byte(swagger.JSONIndexHTML))
				return
			}
			write(w, http.StatusOK, "text/html; charset=utf-8", []byte(swagger.YAMLIndexHTML))
		case "doc.json":
			if !isJSON {
				http.NotFound(w, r)
				return
			}
			docBytes := specBytes
			if requestScheme(r) == swagger.SchemeHTTP {
				docBytes, _ = swagger.PreprocessSchemeInJSONSpec(specBytes, swagger.SchemeHTTP)
			} else {
				docBytes, _ = swagger.PreprocessSchemeInJSONSpec(specBytes, swagger.SchemeHTTPS)
			}
			write(w, http.StatusOK, "application/json; charset=utf-8", docBytes)
		case "doc.yaml":
			if isJSON {
				http.NotFound(w, r)
				return
			}
			docBytes := specBytes
			if requestScheme(r) == swagger.SchemeHTTP {
				docBytes, _ = swagger.PreprocessSchemeInYAMLSpec(specBytes, swagger.SchemeHTTP)
			} else {
				docBytes, _ = swagger.PreprocessSchemeInYAMLSpec(specBytes, swagger.SchemeHTTPS)
			}
			write(w, http.StatusOK, "application/x-yaml; charset=utf-8", docBytes)
		case "swagger-ui.css":
			write(w, http.StatusOK, "text/css; charset=utf-8", []byte(swagger.SwaggerUICSS))
		case "index.css":
			write(w, http.StatusOK, "text/css; charset=utf-8", []byte(swagger.IndexCSS))
		case "swagger-ui-bundle.js":
			write(w, http.StatusOK, "application/javascript; charset=utf-8", []byte(swagger.SwaggerUIBundleJS))
		case "swagger-ui-standalone-preset.js":
			write(w, http.StatusOK, "application/javascript; charset=utf-8", []byte(swagger.SwaggerUIStandalonePresetJS))
		case "favicon-16x16.png":
			write(w, http.StatusOK, "image/png", swagger.Favicon16PNG)
		case "favicon-32x32.png":
			write(w, http.StatusOK, "image/png", swagger.Favicon32PNG)
		default:
			http.NotFound(w, r)
		}
	})
}

func requestScheme(r *http.Request) swagger.Scheme {
	switch r.Header.Get("X-Forwarded-Proto") {
	case "http":
		return swagger.SchemeHTTP
	case "https":
		return swagger.SchemeHTTPS
	}
	if r.TLS != nil {
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

func write(w http.ResponseWriter, status int, contentType string, body []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	_, _ = w.Write(body)
}
