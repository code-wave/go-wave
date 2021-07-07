package repository

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type ChatRepository interface {
	GetChatRoom(clientID, hostID, studyPostID int64) (*entity.ChatRoom, *errors.RestErr)
	SaveChatRoom(clientID, hostID, studyPostID int64) (*entity.ChatRoom, *errors.RestErr)
	GetChatRoomByRoomName(roomName string) (*entity.ChatRoom, *errors.RestErr)
	GetChatRoomByID(id int64) (*entity.ChatRoom, *errors.RestErr)
	SaveChatMessage(msg *entity.ChatMessage) (*entity.ChatMessage, *errors.RestErr)
	GetChatMessages(roomID int64) ([]entity.ChatMessage, *errors.RestErr)
}
