package chat

import "github.com/code-wave/go-wave/domain/entity"

// response data type to front server
type Message struct {
	Sender       string `json:"sender"`
	Message      string `json:"message"`
	MessageType  string `json:"message_type"`
	ChatRoomName string `json:"chat_room_name"`
	CreatedAt    string `json:"created_at"`
}

type Messages []Message

func NewMessage(chatMessage entity.ChatMessage) Message {
	return Message{
		Sender:       chatMessage.Sender,
		Message:      chatMessage.Message,
		MessageType:  chatMessage.MessageType,
		ChatRoomName: chatMessage.ChatRoomName,
		CreatedAt:    chatMessage.CreatedAt,
	}
}
