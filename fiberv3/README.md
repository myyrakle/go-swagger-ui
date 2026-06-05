# Fiber v3 adapter

- Swagger UI server for Fiber v3 users
- Minimum version: Fiber v3.0.0, Go 1.25.0

## Setup

```sh
go get github.com/myyrakle/go-swagger-ui/fiberv3@v0.2.1
```

## With Swagger JSON File

```go
package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	swagger "github.com/myyrakle/go-swagger-ui/fiberv3"
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

	"github.com/gofiber/fiber/v3"
	swagger "github.com/myyrakle/go-swagger-ui/fiberv3"
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
