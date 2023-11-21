package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/FKuiv/LocalChat/pkg/db"
	"github.com/FKuiv/LocalChat/pkg/handlers"
	"github.com/FKuiv/LocalChat/pkg/middleware"
	"github.com/FKuiv/LocalChat/pkg/websocket"
)

func StartHTTPServer() {
	DB := db.Init()
	dbHandler := handlers.New(DB)

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Go server for LocalChat")
	}).Methods(http.MethodGet)

	// Endpoints
	muxRouter.HandleFunc("/ws", websocket.WsHandler)

	muxRouter.HandleFunc("/login", dbHandler.Login).Methods(http.MethodPost)

	// User
	muxRouter.HandleFunc("/user", dbHandler.CreateUser).Methods(http.MethodPost)
	muxRouter.HandleFunc("/users", dbHandler.GetAllUsers).Methods(http.MethodGet)
	muxRouter.HandleFunc("/user/{id}", dbHandler.GetUserById).Methods(http.MethodGet)
	muxRouter.HandleFunc("/user/{id}", dbHandler.UpdateUser).Methods(http.MethodPatch)

	// Group
	muxRouter.HandleFunc("/group", dbHandler.CreateGroup).Methods(http.MethodPost)
	muxRouter.HandleFunc("/groups", dbHandler.GetAllGroups).Methods(http.MethodGet)
	muxRouter.HandleFunc("/group/{id}", dbHandler.GetGroupById).Methods(http.MethodGet)
	muxRouter.HandleFunc("/group/{id}", dbHandler.UpdateGroup).Methods(http.MethodPatch)
	muxRouter.HandleFunc("/group/{id}", dbHandler.DeleteGroup).Methods(http.MethodDelete)

	// Message
	muxRouter.HandleFunc("/message", dbHandler.CreateMessage).Methods(http.MethodPost)
	muxRouter.HandleFunc("/messages", dbHandler.GetAllMessages).Methods(http.MethodGet)
	muxRouter.HandleFunc("/message/{id}", dbHandler.GetMessageById).Methods(http.MethodGet)
	muxRouter.HandleFunc("/message/{id}", dbHandler.UpdateMessage).Methods(http.MethodPatch)
	muxRouter.HandleFunc("/message/{id}", dbHandler.DeleteMessage).Methods(http.MethodDelete)

	handler := cors.Default().Handler(muxRouter)

	log.Println("starting http server at localhost:8000")
	http.ListenAndServe(":8000", middleware.CheckUserSession(middleware.SetHeaders(handler), dbHandler))
}
