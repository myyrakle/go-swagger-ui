package swagger

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	swagger "github.com/myyrakle/go-swagger-ui"
)

// Handler는 임베드된 Swagger UI 리소스를 서빙하는 echo 핸들러입니다.
//
// `/swagger/*` 와 같이 catch-all 라우트로 등록하여 사용합니다.
// 처리 경로:
//   - "" (트레일링 슬래시 접근) → index.html 로 리다이렉트
//   - index.html               → 임베드된 엔트리 HTML
//   - swagger-ui.css, index.css → 임베드된 CSS
//   - swagger-ui-bundle.js, swagger-ui-standalone-preset.js → 임베드된 JS
//
// index.html은 ./doc.json 상대 경로로 스펙을 로드하므로,
// 같은 prefix 하위에 doc.json 응답 핸들러를 별도로 등록해야 합니다.
func Handler(c echo.Context) error {
	file := c.Param("*")
	if file == "" {
		base := strings.TrimSuffix(c.Request().URL.Path, "/")
		return c.Redirect(http.StatusMovedPermanently, base+"/index.html")
	}

	switch file {
	case "index.html":
		return c.HTMLBlob(http.StatusOK, []byte(swagger.IndexHTML))
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
}
