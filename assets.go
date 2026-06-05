// Package swagger embeds Swagger UI v5.32.6 static resources and provides
// handlers that can serve them from an echo router.
package swagger

import (
	_ "embed"
	"encoding/base64"
)

// Favicon16PNG and Favicon32PNG are the decoded bytes of the PNGs declared as base64 in favicon.go.
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

// SwaggerUIVersion is the embedded Swagger UI distribution version.
const SwaggerUIVersion = "5.32.6"

//go:embed dist/swagger-ui.css
var SwaggerUICSS string

//go:embed dist/index.css
var IndexCSS string

//go:embed dist/swagger-ui-bundle.js
var SwaggerUIBundleJS string

//go:embed dist/swagger-ui-standalone-preset.js
var SwaggerUIStandalonePresetJS string

// JSONIndexHTML is the Swagger UI entry HTML that loads embedded assets and uses
// ./doc.json from the same prefix as the spec.
const JSONIndexHTML = `<!DOCTYPE html>
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
        url: "./doc.json" + window.location.search,
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

// YAMLIndexHTML is the Swagger UI entry HTML that loads embedded assets and uses
// ./doc.yaml from the same prefix as the spec.
const YAMLIndexHTML = `<!DOCTYPE html>
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
        url: "./doc.yaml" + window.location.search,
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

// IndexHTML is an alias for JSONIndexHTML for existing users.
const IndexHTML = JSONIndexHTML
