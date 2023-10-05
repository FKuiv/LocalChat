package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FKuiv/LocalChat/pkg/models"
)

func (db dbHandler) registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.UserCreateReq
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	fmt.Println("Received this POST data:", request)

	fmt.Println("\nRegister endpoint", w)
	fmt.Println("\nRegister endpoint request", r)
}
