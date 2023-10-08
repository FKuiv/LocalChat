package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/FKuiv/LocalChat/pkg/db"
	"github.com/FKuiv/LocalChat/pkg/handlers"
)

func StartHTTPServer() {
	DB := db.Init()
	dbHandler := handlers.New(DB)

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Go server for LocalChat")
	}).Methods(http.MethodGet)

	// Endpoints
	muxRouter.HandleFunc("/create_user", dbHandler.CreateUser).Methods(http.MethodPost)
	muxRouter.HandleFunc("/users", dbHandler.GetAllUsers).Methods(http.MethodGet)
	muxRouter.HandleFunc("/user/{id}", dbHandler.GetUserById).Methods(http.MethodGet)
	muxRouter.HandleFunc("/user/{id}", dbHandler.UpdateUser).Methods(http.MethodPatch)

	muxRouter.HandleFunc("/message", dbHandler.SaveMessage).Methods(http.MethodPost)

	handler := cors.Default().Handler(muxRouter)

	log.Println("starting http server at localhost:8000")
	http.ListenAndServe(":8000", handler)
}
