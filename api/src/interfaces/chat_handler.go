package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/code-wave/go-wave/application"
	"github.com/code-wave/go-wave/infrastructure/chat"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

type ChatHandler struct {
	userApp      application.UserAppInterface
	studypostApp application.StudyPostInterface
	chatApp      application.ChatAppInterface
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{}
}

// ServeChatWs: "메시지 보내기"를 누르면 여기로 요청
func (chatHandler *ChatHandler) ServeChatWs(chatServer *chat.ChatServer, w http.ResponseWriter, r *http.Request) {
	var chatReq chat.ChatRequest

	if err := json.NewDecoder(r.Body).Decode(&chatReq); err != nil {
		restErr := errors.NewBadRequestError("invalid json body " + err.Error())
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
	}
	defer r.Body.Close()

	// client의 정보를 가져옴
	user, err := chatHandler.userApp.GetUserByID(chatReq.UserID)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	// websocket 기능 추가
	conn, wsErr := upgrader.Upgrade(w, r, nil)
	if err != nil {
		restErr := errors.NewInternalServerError("ws upgrade error " + wsErr.Error())
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
	}

	// client의 정보를 토대로 ChatUser 객체 생성
	chatClient := chat.NewChatUser(user.ID, user.Name, user.Nickname, conn, chatServer)

	// host user ID 가져옴
	hostUserID, err := chatHandler.studypostApp.GetUserIDByPostID(chatReq.StudyPostID)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	// 채팅룸이 기존에 존재하는지 새로 만들어야하는지 확인
	isRoomExist := true
	chatRoom, err := chatHandler.chatApp.GetChatRoom(chatClient.ID, hostUserID, chatReq.StudyPostID)
	if err != nil {
		if err.Message == errors.ErrNoRows { // 기존 채팅룸이 존재하지 않으므로 새로운 방 만듬
			chatRoom, restErr := chatHandler.chatApp.SaveChatRoom(chatClient.ID, hostUserID, chatReq.StudyPostID)
			if restErr != nil {
				w.WriteHeader(restErr.Status)
				w.Write(restErr.ResponseJSON().([]byte))
				return
			}
			isRoomExist = false
			var messages chat.Messages
			chatRes := chat.ChatResponse{
				ChatRoom:     chatRoom,
				ChatMessages: messages,
			}

			cJSON, err := chatRes.ResponseJSON()
			if err != nil {
				restErr := errors.NewInternalServerError("marshalling error " + err.Error())
				w.WriteHeader(restErr.Status)
				w.Write(restErr.ResponseJSON().([]byte))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(cJSON)
		}

		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	// 채팅룸이 이미 존재하면 기존에 존재하던 채팅룸을 보냄
	if isRoomExist {
		var chatMessages chat.Messages
		chatMessages, err = chatHandler.chatApp.GetChatMessages(chatRoom.ID)
		if err != nil {
			w.WriteHeader(err.Status)
			w.Write(err.ResponseJSON().([]byte))
			return
		}
		chatRes := chat.ChatResponse{
			ChatRoom:     chatRoom,
			ChatMessages: chatMessages,
		}

		cJSON, err := chatRes.ResponseJSON()
		if err != nil {
			restErr := errors.NewInternalServerError("marshalling error " + err.Error())
			w.WriteHeader(restErr.Status)
			w.Write(restErr.ResponseJSON().([]byte))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(cJSON)
	}

}
