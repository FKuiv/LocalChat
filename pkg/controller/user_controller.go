package controller

import (
	"github.com/FKuiv/LocalChat/pkg/models"
	repos "github.com/FKuiv/LocalChat/pkg/repos"
)

type repository interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(userId string) (*models.User, error)
	CreateUser(user models.UserRequest) (*models.User, error)
	DeleteUser(userId string) error
	CreateSession(user models.UserRequest) (*models.Session, error)
	UpdateUser(user models.UserRequest, userId string) (*models.User, error)
}

// UserController contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type UserController struct {
	Service repository
}

func InitUserController(userRepo *repos.UserRepo) *UserController {
	return &UserController{
		Service: userRepo,
	}
}
