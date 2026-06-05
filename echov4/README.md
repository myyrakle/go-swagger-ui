# echo v4 swagger ui

- swagger ui server for echo/v4 user
- minimum version: echo/v4.0.0, golang/v1.16.0

## Setup

```
go get github.com/myyrakle/go-swagger-ui/echov4@v0.2.0
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

func main() {
	e := echo.New()

	data, err := os.ReadFile("docs/swagger.json")
	if err != nil {
		panic(err)
	}

    // auto detect json
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

## with swagger YAML file

```go
package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	swagger "github.com/myyrakle/go-swagger-ui/echov4"
)

func main() {
	e := echo.New()

	data, err := os.ReadFile("docs/swagger.yaml")
	if err != nil {
		panic(err)
	}

    // auto detect
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

## with swaggo (docs.go)

```go
package main

import (
	"log"
	"os"

	// Load the generated docs package
	_ "just_test/docs"

	"github.com/labstack/echo/v4"
	swagger "github.com/myyrakle/go-swagger-ui/echov4"
	"github.com/swaggo/swag"
)

func main() {
	e := echo.New()

	jsonDoc, err := swag.ReadDoc()
	if err != nil {
		panic(err)
	}

	swagger.Serve(e, "/docs", []byte(jsonDoc))

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
