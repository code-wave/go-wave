package entity

type ChatMessage struct {
	ChatRoomID   int64
	ChatRoomName string
	SenderID     int64
	Sender       string
	MessageType  string // 필요 없나? (나가기 등 표시할 때)
	Message      string
	CreatedAt    string
}
