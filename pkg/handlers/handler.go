package handlers

import (
	"github.com/FKuiv/LocalChat/pkg/controller"
)

type Handlers struct {
	UserHandler *userHandler
}

func InitHandlers(cont *controller.Controllers) *Handlers {
	return &Handlers{UserHandler: NewUserHandler(*cont.UserController)}
}
