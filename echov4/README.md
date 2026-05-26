# echo v4 modules

## Setup 

```
go get github.com/myyrakle/go-swagger-ui/echov4@v0.0.2
```

## with swagger JSON file 

```go
package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	swagger "github.com/myyrakle/go-swagger-ui/echov4"
)

func newEchoServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.GET("/*", swagger.Handler)

	e.GET("/doc.json", func(c echo.Context) error {
		data, err := os.ReadFile("doc.json")
		if err != nil {
			return echo.NewHTTPError(500, "failed to read doc.json")
		}
		return c.Blob(200, echo.MIMEApplicationJSONCharsetUTF8, data)
	})

	return e
}
```