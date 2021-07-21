package chat

import (
	"context"
	"encoding/json"
	"log"

	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/persistence"
)

type ChatRoom struct {
	register     chan *ChatUser
	unregister   chan *ChatUser
	broadcast    chan []byte
	roomName     string
	users        map[int64]*ChatUser // key=userID
	redisService *persistence.RedisService
	chatRepo     repository.ChatRepository
}

func NewChatRoom(roomName string, redis *persistence.RedisService, chatRepository repository.ChatRepository) *ChatRoom {
	return &ChatRoom{
		register:     make(chan *ChatUser),
		unregister:   make(chan *ChatUser),
		broadcast:    make(chan []byte),
		roomName:     roomName,
		users:        make(map[int64]*ChatUser),
		redisService: redis,
		chatRepo:     chatRepository,
	}
}

func (c *ChatRoom) RunRoom() {
	log.Println(c.redisService)
	go c.subscribeRoom()

	for {
		select {

		case user := <-c.register:
			c.registerUser(user)

		case user := <-c.unregister:
			c.unregisterUser(user)

		case message := <-c.broadcast:
			// DB에 메시지 저장 후 publish
			c.SaveMessage(message)
			c.publishMessage(message)
		}
	}
}

func (c *ChatRoom) registerUser(user *ChatUser) {
	c.users[user.ID] = user
}

func (c *ChatRoom) unregisterUser(user *ChatUser) {
	if _, ok := c.users[user.ID]; ok {
		delete(c.users, user.ID)
	}
}

// broadcastToUsers: 채팅룸안에 있는 유저들에게 메시지를 보냄
func (c *ChatRoom) broadcastToUsers(message []byte) {
	for _, user := range c.users {
		user.Send <- message
	}
}

func (c *ChatRoom) SaveMessage(message []byte) { // TODO: 메시지 저장할 때 에러 발생시 어떻게?
	var chatMessage Message

	err := json.Unmarshal(message, &chatMessage)
	if err != nil {
		log.Fatal("unmarshal error: ", err.Error())
	}

	// DB에 저장하기 위한 객체 생성
	var savedMessage entity.ChatMessage

	savedMessage.ChatRoomID = chatMessage.ChatRoomID
	savedMessage.ChatRoomName = chatMessage.ChatRoomName
	savedMessage.SenderID = chatMessage.SenderID
	savedMessage.Sender = chatMessage.SenderName
	savedMessage.MessageType = chatMessage.MessageType
	savedMessage.Message = chatMessage.Message
	savedMessage.CreatedAt = chatMessage.CreatedAt

	_, restErr := c.chatRepo.SaveChatMessage(&savedMessage)
	if restErr != nil {
		log.Fatal("save chat message error: ", restErr.Message)
	}
}

func (c *ChatRoom) publishMessage(message []byte) {
	err := c.redisService.RClient.Publish(context.Background(), c.roomName, message).Err()
	if err != nil {
		log.Fatal("redis publish error: ", err.Error())
	}
}

func (c *ChatRoom) subscribeRoom() {

	pubsub := c.redisService.RClient.Subscribe(context.Background(), c.roomName)
	ch := pubsub.Channel()

	for msg := range ch { // publish로 받은 메시지를 채팅룸안에 있는 유저들에게 보냄
		c.broadcastToUsers([]byte(msg.Payload))
	}
}
