package chat

import "github.com/code-wave/go-wave/domain/entity"

type ChatServer struct {
	users      entity.Users
	chatUser   map[int64]*ChatUser
	register   chan *ChatUser
	unregister chan *ChatUser
	rooms      map[string]*ChatRoom
}
