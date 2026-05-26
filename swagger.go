// Package swagger는 Swagger UI v5.32.6 정적 리소스를 임베드하여
// echo 라우터에서 서빙할 수 있는 핸들러를 제공합니다.
package swagger

import (
	_ "embed"
	"encoding/base64"
)

// favicon16PNG, favicon32PNG는 favicon.go에 base64로 선언된 PNG를 디코드한 바이트입니다.
var (
	favicon16PNG = mustDecodeBase64(Favicon16)
	favicon32PNG = mustDecodeBase64(Favicon32)
)

func mustDecodeBase64(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic("swagger: invalid base64 favicon: " + err.Error())
	}
	return b
}

// SwaggerUIVersion은 임베드된 Swagger UI 배포본 버전입니다.
const SwaggerUIVersion = "5.32.6"

//go:embed dist/swagger-ui.css
var SwaggerUICSS string

//go:embed dist/index.css
var IndexCSS string

//go:embed dist/swagger-ui-bundle.js
var SwaggerUIBundleJS string

//go:embed dist/swagger-ui-standalone-preset.js
var SwaggerUIStandalonePresetJS string

// IndexHTML은 임베드된 자산을 로드하고 같은 prefix의 ./doc.json을 스펙으로 사용하는
// Swagger UI 엔트리 HTML입니다.
const IndexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Swagger UI</title>
  <link rel="stylesheet" type="text/css" href="./swagger-ui.css" />
  <link rel="stylesheet" type="text/css" href="./index.css" />
  <link rel="icon" type="image/png" href="./favicon-32x32.png" sizes="32x32" />
  <link rel="icon" type="image/png" href="./favicon-16x16.png" sizes="16x16" />
  <script>
    (function () {
      var origMatchMedia = window.matchMedia.bind(window);
      window.matchMedia = function (query) {
        if (typeof query === "string" && query.indexOf("prefers-color-scheme: dark") !== -1) {
          return {
            matches: false,
            media: query,
            onchange: null,
            addListener: function () {},
            removeListener: function () {},
            addEventListener: function () {},
            removeEventListener: function () {},
            dispatchEvent: function () { return false; },
          };
        }
        return origMatchMedia(query);
      };
    })();
  </script>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="./swagger-ui-bundle.js" charset="UTF-8"></script>
  <script src="./swagger-ui-standalone-preset.js" charset="UTF-8"></script>
  <script>
    window.onload = function() {
      window.ui = SwaggerUIBundle({
        url: "./doc.json",
        dom_id: '#swagger-ui',
        deepLinking: true,
        docExpansion: "list",
        persistAuthorization: false,
        syntaxHighlight: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset,
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl,
        ],
        layout: "StandaloneLayout",
      });
    };
  </script>
</body>
</html>`

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
// func Handler(c echo.Context) error {
// 	file := c.Param("*")
// 	if file == "" {
// 		base := strings.TrimSuffix(c.Request().URL.Path, "/")
// 		return c.Redirect(http.StatusMovedPermanently, base+"/index.html")
// 	}

// 	switch file {
// 	case "index.html":
// 		return c.HTMLBlob(http.StatusOK, []byte(IndexHTML))
// 	case "swagger-ui.css":
// 		return c.Blob(http.StatusOK, "text/css; charset=utf-8", []byte(SwaggerUICSS))
// 	case "index.css":
// 		return c.Blob(http.StatusOK, "text/css; charset=utf-8", []byte(IndexCSS))
// 	case "swagger-ui-bundle.js":
// 		return c.Blob(http.StatusOK, "application/javascript; charset=utf-8", []byte(SwaggerUIBundleJS))
// 	case "swagger-ui-standalone-preset.js":
// 		return c.Blob(http.StatusOK, "application/javascript; charset=utf-8", []byte(SwaggerUIStandalonePresetJS))
// 	case "favicon-16x16.png":
// 		return c.Blob(http.StatusOK, "image/png", favicon16PNG)
// 	case "favicon-32x32.png":
// 		return c.Blob(http.StatusOK, "image/png", favicon32PNG)
// 	}

// 	return echo.ErrNotFound
// }
