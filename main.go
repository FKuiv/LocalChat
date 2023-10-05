package main

import (
	"fmt"

	"github.com/FKuiv/LocalChat/pkg/httpserver"
)

func main() {
	fmt.Println("Hello")

	httpserver.StartHTTPServer()
}
