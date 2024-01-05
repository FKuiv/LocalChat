package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"

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
	result := repo.db.Preload("Users").First(&group, "id = ?", groupId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &utils.CustomError{Message: fmt.Sprintf("Group with ID: %s not found", groupId)}
	}

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error getting user with %s", groupId), result.Error)
		return nil, result.Error
	}

	return &group, nil
}

func (repo *GroupRepo) GetExistingGroupsByUsersAndAdmins(userIds []string, adminIds []string) ([]models.Group, error) {
	var groups []models.Group

	// Retrieve all groups
	result := repo.db.Preload("Users").Find(&groups)
	if result.Error != nil {
		return nil, result.Error
	}

	// Filter groups
	filteredGroups := []models.Group{}
	for _, group := range groups {
		// Check if all adminIds are in group.Admins
		isAdmin := true
		for _, adminId := range adminIds {
			if !utils.SliceContainsStr(group.Admins, adminId) {
				isAdmin = false
				break
			}
		}

		// Check if all userIds are in group.Users
		isUser := true
		for _, userId := range userIds {
			if !utils.ContainsUser(group.Users, userId) {
				isUser = false
				break
			}
		}

		if isAdmin && isUser {
			filteredGroups = append(filteredGroups, group)
		}
	}

	return filteredGroups, nil
}

func (repo *GroupRepo) CreateGroup(groupInfo models.GroupRequest) (*models.Group, error) {

	if !groupInfo.IsDm && groupInfo.Name == "" {
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
	usernames := make(map[string]string)

	for _, userId := range groupInfo.UserIDs {
		var user *models.User
		result := repo.db.First(&user, "id = ?", userId)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("Error finding user", result.Error)
			return nil, &utils.CustomError{Message: fmt.Sprintf("Error finding user: %s", result.Error)}
		}

		usernames[userId] = user.Username
		users = append(users, user)

	}

	newGroup := &models.Group{ID: groupId, Name: groupInfo.Name, Usernames: usernames, Users: users, Admins: groupInfo.Admins, IsDm: groupInfo.IsDm}
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

	if err := repo.DeleteGroupPic(groupId); err != nil {
		return &utils.CustomError{Message: fmt.Sprintf("Error deleting group picture: %s", err)}
	}

	// This also deletes the group
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

	if currentGroup.Usernames == nil {
		usernames := make(map[string]string)
		for _, user := range currentGroup.Users {
			usernames[user.ID] = user.Username
		}
		currentGroup.Usernames = usernames
	}

	if newGroupInfo.Name != "" {
		currentGroup.Name = newGroupInfo.Name
	} else if !currentGroup.IsDm {
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
				currentGroup.Usernames[user.ID] = user.Username
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

func (repo *GroupRepo) GetGroupUserIds(groupId string) ([]string, error) {
	var users []models.User
	result := repo.db.Joins("JOIN user_groups ON users.id = user_groups.user_id").
		Where("user_groups.group_id = ?", groupId).
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	var userIds []string
	for _, user := range users {
		userIds = append(userIds, user.ID)
	}

	return userIds, nil
}

func (repo *GroupRepo) GetAllUserGroups(userId string) ([]models.Group, error) {
	var groups []models.Group
	if err := repo.db.Preload("Users").Joins("JOIN user_groups ON groups.id = user_groups.group_id").
		Where("user_groups.user_id = ?", userId).Find(&groups).Error; err != nil {
		return nil, err
	}

	for i := range groups {
		groups[i].Users = []*models.User{}
		groups[i].Messages = []models.Message{}
	}

	return groups, nil
}

func (repo *GroupRepo) GetAllUserGroupIds(userId string) ([]string, error) {
	groups, err := repo.GetAllUserGroups(userId)
	if err != nil {
		return nil, err
	}

	var groupIds []string
	for _, group := range groups {
		groupIds = append(groupIds, group.ID)
	}

	return groupIds, nil
}

func (repo *GroupRepo) SaveGroupPic(picture multipart.File, pictureInfo *multipart.FileHeader, groupId string) error {
	// Create a buffer to store the header of the file in
	pictureHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := picture.Read(pictureHeader); err != nil {
		return err
	}

	// set position back to start.
	if _, err := picture.Seek(0, 0); err != nil {
		return err
	}

	groupTags := map[string]string{"group_id": groupId}

	_, err := repo.minio.PutObject(context.Background(), utils.MINIO_bucket, utils.GroupProfilePicName(groupId), picture, pictureInfo.Size, minio.PutObjectOptions{ContentType: http.DetectContentType(pictureHeader), UserMetadata: groupTags})

	if err != nil {
		return err
	}

	return nil
}

func (repo *GroupRepo) GetGroupPic(groupId string) (string, error) {
	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	urlExpiration := utils.URL_expiration_time

	// Generates a presigned url which expires in a 1 minutes.
	presignedURL, err := repo.minio.PresignedGetObject(context.Background(), utils.MINIO_bucket, utils.GroupProfilePicName(groupId), urlExpiration, reqParams)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}

func (repo *GroupRepo) DeleteGroupPic(groupId string) error {
	err := repo.minio.RemoveObject(context.Background(), utils.MINIO_bucket, utils.GroupProfilePicName(groupId), minio.RemoveObjectOptions{})

	if err != nil {
		return err
	}

	return nil
}
