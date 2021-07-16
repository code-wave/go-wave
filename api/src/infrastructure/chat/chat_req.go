package chat

type ChatRequest struct {
	UserID      int64 `json:"user_id"`
	StudyPostID int64 `json:"study_post_id"`
}

type WsRequest struct {
	UserID       int64  `json:"user_id"`
	ChatRoomName string `json:"chat_room_name"`
}

type ChatServerRequest struct {
	User         *ChatUser
	ChatRoomName string
}
