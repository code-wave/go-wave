package chat

import "C"
import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

//var (
//	newline = []byte{'\n'}
//	space = []byte{' '}
//)

type ChatUser struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	conn      *websocket.Conn
	Send      chan []byte
	ChatRooms map[string]*ChatRoom
	WsServer  *ChatServer
}

func NewChatUser(id int64, name, nickname string, conn *websocket.Conn, wsServer *ChatServer) *ChatUser {
	return &ChatUser{
		ID:        id,
		Name:      name,
		Nickname:  nickname,
		conn:      conn,
		Send:      make(chan []byte),
		ChatRooms: make(map[string]*ChatRoom),
		WsServer:  wsServer,
	}
}

func (c *ChatUser) ReadPump() {
	defer func() {
		c.disconnect()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	//c.conn.SetReadDeadline(time.Now().Add(pongWait))
	//c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, jsonMessage, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		c.handleNewMessage(jsonMessage)
	}

}

func (c *ChatUser) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send: // ok는 채널이 닫혔는지 열렸는지 여부
			c.conn.SetWriteDeadline(time.Now().Add(writeWait)) // TODO: 없앨지 말지 테스트
			if !ok {
				// The WsServer closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			if err != nil {
				log.Fatal("marshal error: ", err.Error())
			}
			w.Write(message) // message = json type

			//// Attach queued chat messages to the current websocket message.
			//n := len(c.Send)
			//for i := 0; i < n; i++ {
			//	w.Write(newline)
			//	w.Write(<-c.Send)
			//}
			//
			//if err := w.Close(); err != nil {
			//	return
			//}
			//case <-ticker.C:
			//	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			//	if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			//		return
			//	}
		}
	}
}

func (c *ChatUser) disconnect() {
	var chatServerReq ChatServerRequest
	chatServerReq.User = c

	for _, room := range c.ChatRooms { // 사용자가 속한 모든 방과 disconnect
		roomName := room.roomName
		chatServerReq.ChatRoomName = roomName
		c.WsServer.Unregister <- chatServerReq
	}

	for _, room := range c.ChatRooms {
		room.unregister <- c
	}
	close(c.Send)
	c.conn.Close()
}

func (c *ChatUser) handleNewMessage(jsonMessage []byte) {
	var chatMessage Message

	err := json.Unmarshal(jsonMessage, &chatMessage)
	if err != nil { // TODO: 나중에 고루틴 종료 관련 처리
		log.Fatal("unmarshal error: ", err.Error())
	}

	roomName := chatMessage.ChatRoomName
	chatRoom, ok := c.ChatRooms[roomName]
	if !ok {
		log.Fatal("chatroom doesn't exist")
	}

	chatRoom.broadcast <- jsonMessage
}
