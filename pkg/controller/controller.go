package controller

import (
	repos "github.com/FKuiv/LocalChat/pkg/repos"
)

// Controllers contains all the controllers
type Controllers struct {
	userController *UserController
}

// InitControllers returns a new Controllers
func InitControllers(repositories *repos.Repositories) *Controllers {
	return &Controllers{
		userController: InitUserController(repositories.UserRepo),
	}
}
