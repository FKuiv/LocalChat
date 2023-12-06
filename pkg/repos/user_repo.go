package repos

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/FKuiv/LocalChat/pkg/utils"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type UserRepo struct {
	db    *gorm.DB
	minio *minio.Client
}

func NewUserRepo(db *gorm.DB, minio *minio.Client) *UserRepo {
	return &UserRepo{
		db:    db,
		minio: minio,
	}
}

func (repo *UserRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := repo.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repo *UserRepo) GetUserById(userId string) (*models.User, error) {
	var user models.User
	result := repo.db.First(&user, "id = ?", userId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &utils.CustomError{Message: fmt.Sprintf("User with ID: %s not found", userId)}
	}

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error getting user with %s", userId), result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepo) CreateUser(userInfo models.UserRequest) (*models.User, error) {
	if userInfo.Username == "" || userInfo.Password == "" {
		return nil, &utils.CustomError{Message: "Username or password not provided"}
	}

	passwordHash, err := utils.HashPassword(userInfo.Password)
	if err != nil {
		log.Println("Error hashing the password", err)
		return nil, err
	}

	userId, userIdErr := gonanoid.New()

	if userIdErr != nil {
		return nil, userIdErr
	}

	newUser := &models.User{ID: userId, Username: userInfo.Username, Password: passwordHash}
	result := repo.db.Create(newUser)

	// It is a hacky solution but GORM doesn't have an error type to check the unique key constraint so I am checking the substring in the error
	if result.Error != nil && strings.Contains(result.Error.Error(), "(SQLSTATE 23505)") {
		return nil, &utils.CustomError{Message: fmt.Sprintf("Username %s is already taken", newUser.Username)}
	}

	// If there are other errors
	if result.Error != nil {
		return nil, result.Error
	}

	return newUser, nil
}

func (repo *UserRepo) DeleteUser(userId string) error {

	var user models.User
	result := repo.db.Preload("Groups").Preload("Session").Where("id = ?", userId).First(&user)

	if result.Error != nil {
		return &utils.CustomError{Message: fmt.Sprintf("Error finding user and their related contents: %s", result.Error)}
	}

	// Delete all connections to messages
	if err := repo.db.Model(&user).Association("Messages").Clear(); err != nil {
		return &utils.CustomError{Message: fmt.Sprintf("Error removing all user connections to messages: %s", result.Error)}
	}

	for _, group := range user.Groups {
		var newAdmins []string
		for _, adminId := range group.Admins {
			if adminId != userId {
				newAdmins = append(newAdmins, adminId)
			}
		}

		if len(group.Admins) == 1 && group.Admins[0] == userId {
			// Permanent delete
			if err := repo.db.Unscoped().Model(&group).Association("Messages").Unscoped().Clear(); err != nil {
				return &utils.CustomError{Message: fmt.Sprintf("Error deleting all user messages in a group they own: %s", err)}
			}

			if err := repo.db.Model(&group).Association("Users").Clear(); err != nil {
				return &utils.CustomError{Message: fmt.Sprintf("Error removing group user association: %s", err)}
			}

			if err := repo.db.Unscoped().Delete(&group).Error; err != nil {
				return &utils.CustomError{Message: fmt.Sprintf("Error deleting group: %s", err)}
			}

		} else {
			if err := repo.db.Model(&user).Association("Groups").Delete(&group); err != nil {
				return &utils.CustomError{Message: fmt.Sprintf("Error deleting user associations with groups: %s", err)}
			}
		}

		if err := repo.db.Model(&group).Update("Admins", newAdmins).Error; err != nil {
			return &utils.CustomError{Message: fmt.Sprintf("Error updating admins of a group: %s", err)}
		}
	}

	if err := repo.db.Unscoped().Model(&user).Association("Session").Unscoped().Clear(); err != nil {
		return &utils.CustomError{Message: fmt.Sprintf("Error deleting user session: %s", err)}
	}

	if err := repo.db.Delete(&user).Error; err != nil {
		return &utils.CustomError{Message: fmt.Sprintf("Error deleting user: %s", err)}
	}

	return nil
}

func (repo *UserRepo) CreateSession(userInfo models.UserRequest) (*models.Session, error) {

	var currentUser models.User
	repo.db.First(&currentUser, "username = ?", userInfo.Username)

	if !utils.CheckPasswordHash(userInfo.Password, currentUser.Password) {
		return nil, &utils.CustomError{Message: "Wrong password"}
	}

	var existingSession models.Session
	repo.db.First(&existingSession, "user_id = ?", currentUser.ID)
	if existingSession != (models.Session{}) {
		return &existingSession, nil
	}

	sessionId, idErr := gonanoid.New()
	if idErr != nil {
		return nil, &utils.CustomError{Message: fmt.Sprintf("Error creating session ID: %s", idErr)}
	}

	newSession := &models.Session{ID: sessionId, UserID: currentUser.ID, ExpiresAt: time.Now().AddDate(0, 0, 7)}
	result := repo.db.Create(newSession)

	if result.Error != nil {
		return nil, &utils.CustomError{Message: fmt.Sprintf("Error creating session: %s", result.Error)}
	}

	return newSession, nil
}

func (repo *UserRepo) DeleteSession(sessionId string, userId string) error {
	var session models.Session
	result := repo.db.Where("id = ?", sessionId).First(&session)
	if result.Error != nil {
		return result.Error
	}

	if session.UserID != userId {
		return &utils.CustomError{Message: "Forbidden"}
	}

	if err := repo.db.Delete(&session).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepo) UpdateUser(newUserInfo models.UserRequest, userId string) (*models.User, error) {

	var currentUser models.User
	result := repo.db.First(&currentUser, "id = ?", userId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &utils.CustomError{Message: fmt.Sprintf("User with ID: %s not found", userId)}
	}

	if result.Error != nil {
		return nil, &utils.CustomError{Message: fmt.Sprintf("Error finding currentuser: %s", result.Error)}
	}

	if newUserInfo.Username != "" {
		usernameCheck := repo.db.Where("name = ?", newUserInfo.Username).First(&currentUser)

		if usernameCheck.RowsAffected == 1 {
			return nil, &utils.CustomError{Message: "Username already exists"}
		}

		currentUser.Username = newUserInfo.Username
	}

	if newUserInfo.Password != "" {
		passwordHash, err := utils.HashPassword(newUserInfo.Password)

		if err != nil {
			return nil, &utils.CustomError{Message: fmt.Sprintf("Error hashing the password: %s", err)}
		}

		currentUser.Password = passwordHash
	}

	repo.db.Save(&currentUser)
	return &currentUser, nil
}
