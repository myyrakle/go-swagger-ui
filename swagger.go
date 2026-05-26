// Package swagger는 Swagger UI v5.32.6 정적 리소스를 임베드하여
// echo 라우터에서 서빙할 수 있는 핸들러를 제공합니다.
package swagger

import (
	_ "embed"
	"encoding/base64"
)

// Favicon16PNG, Favicon32PNG는 favicon.go에 base64로 선언된 PNG를 디코드한 바이트입니다.
var (
	Favicon16PNG = mustDecodeBase64(Favicon16)
	Favicon32PNG = mustDecodeBase64(Favicon32)
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
