package main

import (
	"fmt"

	"github.com/FKuiv/LocalChat/httpserver"
)

func main() {
	fmt.Println("Hello")

	httpserver.StartHTTPServer()
}
