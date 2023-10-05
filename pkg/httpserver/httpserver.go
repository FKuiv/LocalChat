package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/FKuiv/LocalChat/pkg/db"
)

func StartHTTPServer() {
	DB := db.Init()
	dbHandler := New(DB)

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	}).Methods(http.MethodGet)

	// Endpoints
	muxRouter.HandleFunc("/register", dbHandler.registerHandler).Methods(http.MethodPost)

	handler := cors.Default().Handler(muxRouter)

	fmt.Println("starting http server at localhost:8000")
	http.ListenAndServe(":8000", handler)
}
