package echov4swagger

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestServeWithPrefixWithoutTrailingSlash(t *testing.T) {
	e := echo.New()
	Serve(e, "/docs", []byte(`{"swagger":"2.0","schemes":["http","https"]}`))

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
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			if rec.Code != tt.want {
				t.Fatalf("status = %d, want %d; body = %q", rec.Code, tt.want, rec.Body.String())
			}
		})
	}
}
