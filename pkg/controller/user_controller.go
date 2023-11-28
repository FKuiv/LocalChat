package controller

import (
	"github.com/FKuiv/LocalChat/pkg/models"
	repos "github.com/FKuiv/LocalChat/pkg/repos"
)

type repository interface {
	GetExistingUser(username string) *models.User
	CreateUser(user models.User) (*models.User, error)
	GetUsers() ([]models.User, error)
}

// UserController contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type UserController struct {
	service repository
}

func InitUserController(userRepo *repos.UserRepo) *UserController {
	return &UserController{
		service: userRepo,
	}
}
