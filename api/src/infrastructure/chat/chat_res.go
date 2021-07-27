package chat

import (
	"encoding/json"

	"github.com/code-wave/go-wave/domain/entity"
)

type ChatResponse struct {
	ChatRoom     *entity.ChatRoom `json:"chat_room"`
	ChatMessages Messages         `json:"chat_messages"` // 기존에 저장된 메시지들
}

func (r *ChatResponse) ResponseJSON() ([]byte, error) {

	chatJson, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return chatJson, nil
}
