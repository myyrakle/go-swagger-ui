# Fiber v2 adapter

- Swagger UI server for Fiber v2 users
- Minimum version: Fiber v2.0.0, Go 1.14

## Setup

```sh
go get github.com/myyrakle/go-swagger-ui/fiberv2@v0.1.10
```

## With Swagger JSON File

```go
package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/myyrakle/go-swagger-ui/fiberv2"
)

func main() {
	app := fiber.New()

	data, err := os.ReadFile("docs/swagger.json")
	if err != nil {
		panic(err)
	}

	swagger.Serve(app, "/docs", data)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":28080"
	}

	log.Printf("fiber server listening on %s", addr)
	if err := app.Listen(addr); err != nil {
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

	"github.com/gofiber/fiber/v2"
	swagger "github.com/myyrakle/go-swagger-ui/fiberv2"
)

func main() {
	app := fiber.New()

	data, err := os.ReadFile("docs/swagger.yaml")
	if err != nil {
		panic(err)
	}

	swagger.Serve(app, "/docs", data)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":28080"
	}

	log.Printf("fiber server listening on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
```
