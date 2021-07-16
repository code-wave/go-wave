package chat

import (
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/persistence"
)

type ChatServer struct {
	//users      entity.Users
	chatUsers    map[int64]*ChatUser // key=userID
	Register     chan ChatServerRequest
	Unregister   chan ChatServerRequest
	rooms        map[string]*ChatRoom // key=ChatRoomName
	redisService *persistence.RedisService
	chatRepo     repository.ChatRepository
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		chatUsers:  make(map[int64]*ChatUser),
		Register:   make(chan ChatServerRequest),
		Unregister: make(chan ChatServerRequest),
		rooms:      make(map[string]*ChatRoom),
	}
}

func (c *ChatServer) Run() {
	for {
		select {

		case req := <-c.Register:
			c.registerUser(req.User, req.ChatRoomName)

		case req := <-c.Unregister:
			c.unregisterUser(req.User, req.ChatRoomName)
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

//CreateRoom: 메모리상에 채팅룸 생성
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
