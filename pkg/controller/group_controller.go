package controller

import (
	"mime/multipart"

	"github.com/FKuiv/LocalChat/pkg/models"
	repos "github.com/FKuiv/LocalChat/pkg/repository"
)

type group_repository interface {
	GetAllGroups() ([]models.Group, error)
	GetGroupById(groupId string) (*models.Group, error)
	GetExistingGroupsByUsersAndAdmins(userIds []string, adminIds []string) ([]models.Group, error)
	CreateGroup(group models.GroupRequest) (*models.Group, error)
	DeleteGroup(groupId string, userId string) error
	UpdateGroup(group models.GroupRequest, groupId string) (*models.Group, error)
	GetGroupUserIds(groupId string) ([]string, error)
	GetAllUserGroups(userId string) ([]models.Group, error)
	GetAllUserGroupIds(userId string) ([]string, error)
	SaveGroupPic(picture multipart.File, pictureInfo *multipart.FileHeader, groupId string) error
	GetGroupPic(groupId string) (string, error)
	DeleteGroupPic(groupId string) error
}

type GroupController struct {
	Service group_repository
}

func InitGroupController(groupRepo *repos.GroupRepo) *GroupController {
	return &GroupController{
		Service: groupRepo,
	}
}
