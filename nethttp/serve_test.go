package nethttpswagger

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerWithPrefixWithoutTrailingSlash(t *testing.T) {
	handler := Handler("/docs", []byte(`{"swagger":"2.0","schemes":["http","https"]}`))

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

			handler.ServeHTTP(rec, req)

			if rec.Code != tt.want {
				t.Fatalf("status = %d, want %d; body = %q", rec.Code, tt.want, rec.Body.String())
			}
		})
	}
}
