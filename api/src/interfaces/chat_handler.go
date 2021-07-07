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

func (chatHandler *ChatHandler) ServeChatWs(chatServer *chat.ChatServer, w http.ResponseWriter, r *http.Request) {
	var chatReq chat.ChatRequest

	if err := json.NewDecoder(r.Body).Decode(&chatReq); err != nil {
		restErr := errors.NewBadRequestError("invalid json body " + err.Error())
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
	}
	defer r.Body.Close()

	user, err := chatHandler.userApp.GetUserByID(chatReq.UserID)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	conn, wsErr := upgrader.Upgrade(w, r, nil)
	if err != nil {
		restErr := errors.NewInternalServerError("ws upgrade error " + wsErr.Error())
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
	}

	chatClient := chat.NewChatUser(user.ID, user.Name, user.Nickname, conn)

	hostUserID, err := chatHandler.studypostApp.GetUserIDByPostID(chatReq.StudyPostID)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	isRoomExist := true
	chatRoom, err := chatHandler.chatApp.GetChatRoom(chatClient.ID, hostUserID, chatReq.StudyPostID)
	if err != nil {
		if err.Message == errors.ErrNoRows {
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

	//get message from db by chatRoomID
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

	//go readpump go wrtepump

	//

}
