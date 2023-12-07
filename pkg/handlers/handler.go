package handlers

import (
	"github.com/FKuiv/LocalChat/pkg/controller"
)

type Handlers struct {
	UserHandler    *userHandler
	GroupHandler   *groupHandler
	MessageHandler *messageHandler
	WsHandler      *wsHandler
}

func InitHandlers(cont *controller.Controllers) *Handlers {
	return &Handlers{
		UserHandler:    NewUserHandler(*cont.UserController),
		GroupHandler:   NewGroupHandler(*cont.GroupController),
		MessageHandler: NewMessageHandler(*cont.MessageController),
		WsHandler:      NewWsHandler(*cont),
	}
}
