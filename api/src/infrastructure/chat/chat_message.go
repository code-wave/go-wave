package chat

import "github.com/code-wave/go-wave/domain/entity"

// response data type to front server
type Message struct {
	ChatRoomID   int64  `json:"chat_room_id"`
	ChatRoomName string `json:"chat_room_name"`
	SenderID     int64  `json:"sender_id"`
	SenderName   string `json:"sender_name"`
	Message      string `json:"message"`
	MessageType  string `json:"message_type"`
	CreatedAt    string `json:"created_at"`
}

type Messages []Message

func NewMessage(chatMessage entity.ChatMessage) Message {
	return Message{
		ChatRoomID:   chatMessage.ChatRoomID,
		ChatRoomName: chatMessage.ChatRoomName,
		SenderID:     chatMessage.SenderID,
		SenderName:   chatMessage.Sender,
		Message:      chatMessage.Message,
		MessageType:  chatMessage.MessageType,
		CreatedAt:    chatMessage.CreatedAt,
	}
}
