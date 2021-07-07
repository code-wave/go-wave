package chat

type ChatRequest struct {
	UserID      int64 `json:"user_id"`
	StudyPostID int64 `json:"study_post_id"`
}
