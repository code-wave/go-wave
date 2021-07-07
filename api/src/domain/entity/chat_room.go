package entity

type ChatRoom struct {
	ID          int64  `json:"id"`
	RoomName    string `json:"room_name"`
	ClientID    int64  `json:"client_id"` // 메시지 보내기 누른 사람 ID
	HostID      int64  `json:"host_id"`   // 게시글 쓴 사람 ID
	StudyPostID int64  `json:"study_post_id"`
}
