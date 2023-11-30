package controller

import (
	"github.com/FKuiv/LocalChat/pkg/models"
	repos "github.com/FKuiv/LocalChat/pkg/repos"
)

type message_repository interface {
	GetAllMessages() ([]models.Message, error)
	GetMessageById(messageId string) (*models.Message, error)
	CreateMessage(message models.MessageRequest, userId string) (*models.Message, error)
	DeleteMessage(messageId string, userId string) error
	UpdateMessage(newMessage models.UpdateMessage, messageId string) (*models.Message, error)
}

type MessageController struct {
	Service message_repository
}

func InitMessageController(messageRepo *repos.MessageRepo) *MessageController {
	return &MessageController{
		Service: messageRepo,
	}
}
