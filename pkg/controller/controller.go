package controller

import (
	repos "github.com/FKuiv/LocalChat/pkg/repos"
)

// Controllers contains all the controllers
type Controllers struct {
	UserController    *UserController
	GroupController   *GroupController
	MessageController *MessageController
}

// InitControllers returns a new Controllers
func InitControllers(repositories *repos.Repositories) *Controllers {
	return &Controllers{
		UserController:    InitUserController(repositories.UserRepo),
		GroupController:   InitGroupController(repositories.GroupRepo),
		MessageController: InitMessageController(repositories.MessageRepo),
	}
}
