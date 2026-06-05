# gin swagger ui

- Swagger UI server for Gin users
- Minimum version: Gin v1.10.1, Go 1.20

## Setup

```sh
go get github.com/myyrakle/go-swagger-ui/ginv1@v0.2.1
```

## With Swagger JSON File

```go
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swagger "github.com/myyrakle/go-swagger-ui/ginv1"
)

func main() {
	r := gin.Default()

	data, err := os.ReadFile("docs/swagger.json")
	if err != nil {
		panic(err)
	}

	swagger.Serve(r, "/docs", data)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":28080"
	}

	log.Printf("gin server listening on %s", addr)
	if err := r.Run(addr); err != nil {
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

	"github.com/gin-gonic/gin"
	swagger "github.com/myyrakle/go-swagger-ui/gin"
)

func main() {
	r := gin.Default()

	data, err := os.ReadFile("docs/swagger.yaml")
	if err != nil {
		panic(err)
	}

	swagger.Serve(r, "/docs", data)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":28080"
	}

	log.Printf("gin server listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
```
