package application

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/chat"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

var _ ChatAppInterface = &ChatApp{}

type ChatApp struct {
	chatRepo repository.ChatRepository
}

type ChatAppInterface interface {
	GetChatRoom(clientID, hostID, studyPostID int64) (*entity.ChatRoom, *errors.RestErr)
	SaveChatRoom(clientID, hostID, studyPostID int64) (*entity.ChatRoom, *errors.RestErr)
	GetChatRoomByRoomName(roomName string) (*entity.ChatRoom, *errors.RestErr)
	GetChatRoomByID(id int64) (*entity.ChatRoom, *errors.RestErr)
	SaveChatMessage(msg *entity.ChatMessage) (*entity.ChatMessage, *errors.RestErr)
	GetChatMessages(roomID int64) (chat.Messages, *errors.RestErr)
}

func NewChatApp(chatRepo repository.ChatRepository) *ChatApp {
	return &ChatApp{
		chatRepo: chatRepo,
	}
}

func (chatApp *ChatApp) GetChatRoom(clientID, hostID, studyPostID int64) (*entity.ChatRoom, *errors.RestErr) {
	return chatApp.chatRepo.GetChatRoom(clientID, hostID, studyPostID)
}
func (chatApp *ChatApp) SaveChatRoom(clientID, hostID, studyPostID int64) (*entity.ChatRoom, *errors.RestErr) {
	return chatApp.chatRepo.SaveChatRoom(clientID, hostID, studyPostID)
}
func (chatApp *ChatApp) GetChatRoomByRoomName(roomName string) (*entity.ChatRoom, *errors.RestErr) {
	return nil, nil
}
func (chatApp *ChatApp) GetChatRoomByID(id int64) (*entity.ChatRoom, *errors.RestErr) {
	return nil, nil
}
func (chatApp *ChatApp) SaveChatMessage(msg *entity.ChatMessage) (*entity.ChatMessage, *errors.RestErr) {
	return nil, nil
}

func (chatApp *ChatApp) GetChatMessages(roomID int64) (chat.Messages, *errors.RestErr) {

	chatMessages, err := chatApp.chatRepo.GetChatMessages(roomID)
	if err != nil {
		return nil, err
	}

	var messages chat.Messages

	for _, msg := range chatMessages {
		messages = append(messages, chat.NewMessage(msg))
	}

	return messages, nil
}
