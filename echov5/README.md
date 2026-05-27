# echo v5 swagger ui

- Swagger UI server for Echo v5 users
- Minimum version: Echo v5.0.0, Go 1.25.0

## Setup

```sh
go get github.com/myyrakle/go-swagger-ui/echov5@v0.1.10
```

## With Swagger JSON File

```go
package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v5"
	swagger "github.com/myyrakle/go-swagger-ui/echov5"
)

func main() {
	e := echo.New()

	data, err := os.ReadFile("docs/swagger.json")
	if err != nil {
		panic(err)
	}

	swagger.Serve(e, "/docs", data)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":28080"
	}

	log.Printf("echo server listening on %s", addr)
	if err := e.Start(addr); err != nil {
		log.Fatal(err)
	}
}
```

## With Swagger YAML File

```go
package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v5"
	swagger "github.com/myyrakle/go-swagger-ui/echov5"
)

func main() {
	e := echo.New()

	data, err := os.ReadFile("docs/swagger.yaml")
	if err != nil {
		panic(err)
	}

	swagger.Serve(e, "/docs", data)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":28080"
	}

	log.Printf("echo server listening on %s", addr)
	if err := e.Start(addr); err != nil {
		log.Fatal(err)
	}
}
```
