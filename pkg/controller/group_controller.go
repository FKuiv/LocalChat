package controller

import (
	"github.com/FKuiv/LocalChat/pkg/models"
	repos "github.com/FKuiv/LocalChat/pkg/repository"
)

type group_repository interface {
	GetAllGroups() ([]models.Group, error)
	GetGroupById(groupId string) (*models.Group, error)
	CreateGroup(group models.GroupRequest) (*models.Group, error)
	DeleteGroup(groupId string, userId string) error
	UpdateGroup(group models.GroupRequest, groupId string) (*models.Group, error)
}

type GroupController struct {
	Service group_repository
}

func InitGroupController(groupRepo *repos.GroupRepo) *GroupController {
	return &GroupController{
		Service: groupRepo,
	}
}
