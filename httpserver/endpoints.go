package httpserver

import (
	"fmt"
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("\nRegister endpoint", w)
	fmt.Println("\nRegister endpoint request", r)
}
