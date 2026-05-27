package fiberv3swagger

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestServeWithPrefixWithoutTrailingSlash(t *testing.T) {
	app := fiber.New()
	Serve(app, "/docs", []byte(`{"swagger":"2.0","schemes":["http","https"]}`))

	tests := []struct {
		path string
		want int
	}{
		{path: "/docs", want: http.StatusMovedPermanently},
		{path: "/docs/", want: http.StatusOK},
		{path: "/docs/index.html", want: http.StatusOK},
		{path: "/docs/doc.json", want: http.StatusOK},
		{path: "/docs/swagger-ui.css", want: http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "http://example.com"+tt.path, nil)
			if err != nil {
				t.Fatalf("new request: %v", err)
			}

			res, err := app.Test(req)
			if err != nil {
				t.Fatalf("app test: %v", err)
			}
			defer res.Body.Close()

			if res.StatusCode != tt.want {
				t.Fatalf("status = %d, want %d", res.StatusCode, tt.want)
			}
		})
	}
}
