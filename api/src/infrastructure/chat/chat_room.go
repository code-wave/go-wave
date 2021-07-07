package chat

type ChatRoom struct {
	register   chan (*ChatUser)
	unregister chan (*ChatUser)
	broadcast  chan (*Message)
	roomName   string
	users      map[int64]*ChatUser
}
