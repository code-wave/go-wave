package chat

import (
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/persistence"
)

type ChatServer struct {
	//users      entity.Users
	chatUsers    map[int64]*ChatUser // key=userID
	register     chan ChatServerRequest
	unregister   chan ChatServerRequest
	rooms        map[string]*ChatRoom // key=roomName
	redisService *persistence.RedisService
	chatRepo     repository.ChatRepository
}

type ChatServerRequest struct {
	user     *ChatUser
	roomName string
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		chatUsers:  make(map[int64]*ChatUser),
		register:   make(chan ChatServerRequest),
		unregister: make(chan ChatServerRequest),
		rooms:      make(map[string]*ChatRoom),
	}
}

func (c *ChatServer) Run() {
	for {
		select {

		case req := <-c.register:
			c.registerUser(req.user, req.roomName)

		case req := <-c.unregister:
			c.unregisterUser(req.user, req.roomName)
		}
	}
}

func (c *ChatServer) registerUser(user *ChatUser, roomName string) {
	c.chatUsers[user.ID] = user
	c.rooms[roomName].register <- user
}

func (c *ChatServer) unregisterUser(user *ChatUser, roomName string) {
	if _, ok := c.chatUsers[user.ID]; ok {
		delete(c.chatUsers, user.ID)
		c.rooms[roomName].unregister <- user
	}
}

func (c *ChatServer) CreateRoom(roomName string) *ChatRoom {
	room := NewChatRoom(roomName, c.redisService, c.chatRepo)
	c.rooms[roomName] = room

	go room.RunRoom()

	return room
}

func (c *ChatServer) GetRoomByName(roomName string) *ChatRoom {
	if room, ok := c.rooms[roomName]; ok {
		return room
	}
	return nil
}
