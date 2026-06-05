# net/http adapter

- Swagger UI server for Go's standard net/http package
- Minimum version: Go 1.18

## Setup

```sh
go get github.com/myyrakle/go-swagger-ui/nethttp@v0.2.0
```

## With Swagger JSON File

```go
package main

import (
	"log"
	"net/http"
	"os"

	swagger "github.com/myyrakle/go-swagger-ui/nethttp"
)

func main() {
	data, err := os.ReadFile("docs/swagger.json")
	if err != nil {
		panic(err)
	}

	handler := swagger.Handler("/docs", data)
	http.Handle("/docs", handler)
	http.Handle("/docs/", handler)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":28080"
	}

	log.Printf("http server listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
```

## With Swagger YAML File

```go
package main

import (
	"log"
	"net/http"
	"os"

	swagger "github.com/myyrakle/go-swagger-ui/nethttp"
)

func main() {
	data, err := os.ReadFile("docs/swagger.yaml")
	if err != nil {
		panic(err)
	}

	handler := swagger.Handler("/docs", data)
	http.Handle("/docs", handler)
	http.Handle("/docs/", handler)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":28080"
	}

	log.Printf("http server listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
```
