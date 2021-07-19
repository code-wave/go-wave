package chat

import "github.com/gorilla/websocket"

type ChatUser struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	conn      *websocket.Conn
	Send      chan ([]byte)
	ChatRooms map[string]*ChatRoom
}

func NewChatUser(id int64, name, nickname string, conn *websocket.Conn) *ChatUser {
	return &ChatUser{
		ID:        id,
		Name:      name,
		Nickname:  nickname,
		conn:      conn,
		Send:      make(chan []byte),
		ChatRooms: make(map[string]*ChatRoom),
	}
}
