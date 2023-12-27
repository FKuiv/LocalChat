package repository

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

type MessageRepo struct {
	db    *gorm.DB
	minio *minio.Client
}

func NewMessageRepo(db *gorm.DB, minio *minio.Client) *MessageRepo {
	return &MessageRepo{
		db:    db,
		minio: minio,
	}
}

func (repo *MessageRepo) GetAllMessages() ([]models.Message, error) {
	var messages []models.Message
	result := repo.db.Find(&messages)

	if result.Error != nil {
		log.Println("Error getting the messsages", result.Error)
		return nil, result.Error
	}

	return messages, nil
}

func (repo *MessageRepo) GetMessageById(messageId string) (*models.Message, error) {
	var message models.Message
	result := repo.db.First(&message, "id = ?", messageId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &utils.CustomError{Message: fmt.Sprintf("Message with ID: %s not found", messageId)}
	}

	if result.Error != nil {
		log.Println("Error getting the message", result.Error)
		return nil, result.Error
	}

	return &message, nil
}

func (repo *MessageRepo) CreateMessage(message models.MessageRequest, userId string) (*models.Message, error) {
	messageId, messageIdErr := gonanoid.New()

	if messageIdErr != nil {
		log.Println("Error creating message ID", messageIdErr)
		return nil, messageIdErr
	}

	newMessage := &models.Message{ID: messageId, Content: message.Content, UserID: userId, GroupID: message.GroupID}
	result := repo.db.Create(newMessage)

	if result.Error != nil {
		log.Println("error saving message:", result.Error)
		return nil, result.Error
	}

	return newMessage, nil
}

func (repo *MessageRepo) DeleteMessage(messageId string, userId string) error {
	message, err := repo.GetMessageById(messageId)

	if err != nil {
		return err
	}

	if message.UserID != userId && message.UserID != "" {
		return &utils.CustomError{Message: "User does not own this message, therefore cannot delete it"}
	}

	// Attempt to delete the message
	if err := repo.db.Delete(&message).Error; err != nil {
		fmt.Println("Error deleting message:", err)
		return err
	}

	return nil
}

func (repo *MessageRepo) UpdateMessage(newMessageInfo models.UpdateMessage, messageId string) (*models.Message, error) {
	currentMessage, err := repo.GetMessageById(messageId)

	if err != nil {
		log.Println("Error getting the message", err)
		return nil, err
	}

	if newMessageInfo.Content != "" {
		currentMessage.Content = newMessageInfo.Content
	} else {
		return nil, &utils.CustomError{Message: "Message cannot have empty content. If you want to delete it, then use delete."}
	}

	repo.db.Save(&currentMessage)
	return currentMessage, nil
}

func (repo *MessageRepo) GetMessagesByGroup(groupId string, messageCount int) ([]models.Message, error) {
	var messages []models.Message
	result := repo.db.Limit(messageCount).Where("group_id = ?", groupId).Order("created_at ASC").Find(&messages)

	if result.Error != nil {
		log.Println("Error getting messages by group:", result.Error)
		return nil, result.Error
	}

	return messages, nil
}
