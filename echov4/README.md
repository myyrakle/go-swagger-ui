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