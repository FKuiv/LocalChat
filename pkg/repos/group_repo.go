package repos

import (
	"errors"
	"fmt"
	"log"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type GroupRepo struct {
	db    *gorm.DB
	minio *minio.Client
}

func NewGroupRepo(db *gorm.DB, minio *minio.Client) *GroupRepo {
	return &GroupRepo{
		db:    db,
		minio: minio,
	}
}

func (repo *GroupRepo) GetAllGroups() ([]models.Group, error) {
	var groups []models.Group
	result := repo.db.Find(&groups)

	if result.Error != nil {
		return nil, result.Error
	}

	return groups, nil
}

func (repo *GroupRepo) GetGroupById(groupId string) (*models.Group, error) {
	var group models.Group
	result := repo.db.First(&group, "id = ?", groupId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &utils.CustomError{Message: fmt.Sprintf("Group with ID: %s not found", groupId)}
	}

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error getting user with %s", groupId), result.Error)
		return nil, result.Error
	}

	return &group, nil
}

func (repo *GroupRepo) CreateGroup(groupInfo models.GroupRequest) (*models.Group, error) {
	if groupInfo.Name == "" {
		return nil, &utils.CustomError{Message: "Group name can't be empty"}
	}

	if len(groupInfo.Admins) == 0 || len(groupInfo.UserIDs) == 0 {
		return nil, &utils.CustomError{Message: "There needs to be at least 1 admin and user in group"}
	}

	groupId, groupIdErr := gonanoid.New()

	if groupIdErr != nil {
		return nil, &utils.CustomError{Message: fmt.Sprintf("Error creating group ID: %s", groupIdErr)}
	}

	var users []*models.User

	for _, userId := range groupInfo.UserIDs {
		var user *models.User
		result := repo.db.First(&user, "id = ?", userId)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("Error finding user", result.Error)
			return nil, &utils.CustomError{Message: fmt.Sprintf("Error finding user: %s", result.Error)}
		} else {
			users = append(users, user)
		}

	}

	newGroup := &models.Group{ID: groupId, Name: groupInfo.Name, Users: users, Admins: groupInfo.Admins, IsDm: groupInfo.IsDm}
	result := repo.db.Create(newGroup)

	if result.Error != nil {
		return nil, &utils.CustomError{Message: fmt.Sprintf("Error creating group: %s", result.Error)}
	}

	return newGroup, nil
}

func (repo *GroupRepo) DeleteGroup(groupId string, userId string) error {

	var group models.Group
	result := repo.db.Where("id = ?", groupId).First(&group)

	if result.Error != nil {
		return &utils.CustomError{Message: fmt.Sprintf("Couldn't find group with ID %s. Error: %d", groupId, result.Error)}
	}

	isAdmin := false
	for _, adminId := range group.Admins {
		if userId == adminId {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		return &utils.CustomError{Message: "User needs to be admin to delete this group"}
	}

	// Need to delete all the messages inside a group first
	if err := repo.db.Unscoped().Model(&group).Association("Messages").Unscoped().Clear(); err != nil {
		fmt.Println("Error deleting all the messages in a group", err)
		return &utils.CustomError{Message: fmt.Sprintf("Failed to delete all group messages: %s", err)}
	}

	if err := repo.db.Select("Users").Delete(&group).Error; err != nil {
		fmt.Println("Error removing references from user_groups table", err)
		return &utils.CustomError{Message: fmt.Sprintf("Failed to remove references from user_groups table: %s", err)}
	}

	return nil
}

func (repo *GroupRepo) UpdateGroup(newGroupInfo models.GroupRequest, groupId string) (*models.Group, error) {
	var currentGroup models.Group
	result := repo.db.Where("id = ?", groupId).First(&currentGroup)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &utils.CustomError{Message: fmt.Sprintf("Group with ID: %s not found", groupId)}
	}

	if result.Error != nil {
		log.Println("Error getting the group", result.Error)
		return nil, &utils.CustomError{Message: fmt.Sprintf("Error getting the group: %s", result.Error)}
	}

	if newGroupInfo.Name != "" {
		currentGroup.Name = newGroupInfo.Name
	} else {
		return nil, &utils.CustomError{Message: "Group name cannot be empty"}
	}

	if len(newGroupInfo.UserIDs) != 0 {
		var users []*models.User

		for _, userId := range newGroupInfo.UserIDs {
			var user *models.User
			result := repo.db.First(&user, "id = ?", userId)

			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				log.Println("Error finding user", result.Error)
				return nil, &utils.CustomError{Message: fmt.Sprintf("Error finding user: %s", result.Error)}
			} else {
				users = append(users, user)
			}

		}

		currentGroup.Users = users
	} else {
		return nil, &utils.CustomError{Message: "A group cannot have 0 users"}
	}

	if len(newGroupInfo.Admins) != 0 {
		currentGroup.Admins = newGroupInfo.Admins
	} else {
		return nil, &utils.CustomError{Message: "A group cannot have 0 admins"}
	}

	repo.db.Save(&currentGroup)

	return &currentGroup, nil
}
