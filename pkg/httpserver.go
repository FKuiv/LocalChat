package pkg

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/FKuiv/LocalChat/pkg/controller"
	"github.com/FKuiv/LocalChat/pkg/db"
	"github.com/FKuiv/LocalChat/pkg/handlers"
	"github.com/FKuiv/LocalChat/pkg/middleware"
	"github.com/FKuiv/LocalChat/pkg/repository"
	"github.com/FKuiv/LocalChat/pkg/websocket"
)

func StartHTTPServer() {
	dbconn := db.Init()
	minioConn := db.InitMinio()

	repositories := repository.InitRepositories(dbconn.GetDB(), minioConn.GetMinio())
	controllers := controller.InitControllers(repositories)
	handlers := handlers.InitHandlers(controllers)

	hub := websocket.NewHub(controllers)
	go hub.Run()

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
	// Endpoints
	muxRouter.HandleFunc("/ws", hub.Handle)
	muxRouter.HandleFunc("/ws/refresh", hub.RefreshWs).Methods(http.MethodPost)

	muxRouter.HandleFunc("/login", handlers.UserHandler.Login).Methods(http.MethodPost)
	muxRouter.HandleFunc("/logout", handlers.UserHandler.Logout).Methods(http.MethodGet)

	muxRouter.HandleFunc("/profilepic", handlers.UserHandler.UploadProfilePic).Methods(http.MethodPost)
	muxRouter.HandleFunc("/profilepic", handlers.UserHandler.GetProfilePic).Methods(http.MethodGet)
	// User
	muxRouter.HandleFunc("/user", handlers.UserHandler.CreateUser).Methods(http.MethodPost)
	muxRouter.HandleFunc("/users", handlers.UserHandler.GetAllUsers).Methods(http.MethodGet)
	muxRouter.HandleFunc("/user/{id}", handlers.UserHandler.GetUserById).Methods(http.MethodGet)
	muxRouter.HandleFunc("/user", handlers.UserHandler.UpdateUser).Methods(http.MethodPatch)
	muxRouter.HandleFunc("/user_delete", handlers.UserHandler.DeleteUser).Methods(http.MethodDelete)

	// Group
	muxRouter.HandleFunc("/group", handlers.GroupHandler.CreateGroup).Methods(http.MethodPost)
	muxRouter.HandleFunc("/groups", handlers.GroupHandler.GetAllGroups).Methods(http.MethodGet)
	muxRouter.HandleFunc("/group/{id}", handlers.GroupHandler.GetGroupById).Methods(http.MethodGet)
	muxRouter.HandleFunc("/group/{id}", handlers.GroupHandler.UpdateGroup).Methods(http.MethodPatch)
	muxRouter.HandleFunc("/group/{id}", handlers.GroupHandler.DeleteGroup).Methods(http.MethodDelete)

	// Message
	muxRouter.HandleFunc("/message", handlers.MessageHandler.CreateMessage).Methods(http.MethodPost)
	muxRouter.HandleFunc("/messages", handlers.MessageHandler.GetAllMessages).Methods(http.MethodGet)
	muxRouter.HandleFunc("/message/{id}", handlers.MessageHandler.GetMessageById).Methods(http.MethodGet)
	muxRouter.HandleFunc("/message/{id}", handlers.MessageHandler.UpdateMessage).Methods(http.MethodPatch)
	muxRouter.HandleFunc("/message/{id}", handlers.MessageHandler.DeleteMessage).Methods(http.MethodDelete)

	corsInstance := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:8080"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := corsInstance.Handler(muxRouter)

	log.Println("starting http server at localhost:8000")
	http.ListenAndServe(":8000", middleware.CheckUserSession(middleware.SetHeaders(handler), dbconn.GetDB()))
}
