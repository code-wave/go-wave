package interfaces

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/code-wave/go-wave/infrastructure/helpers"

	"github.com/code-wave/go-wave/application"
	"github.com/code-wave/go-wave/infrastructure/chat"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatHandler struct {
	userApp      application.UserAppInterface
	studyPostApp application.StudyPostInterface
	chatApp      application.ChatAppInterface
}

func NewChatHandler(userApp application.UserAppInterface, studyPostApp application.StudyPostInterface, chatApp application.ChatAppInterface) *ChatHandler {
	return &ChatHandler{
		userApp:      userApp,
		studyPostApp: studyPostApp,
		chatApp:      chatApp,
	}
}

// ServeChatWs: roomName과 유저정보를 보내면 websocket 연결시켜줌
func (chatHandler *ChatHandler) ServeChatWs(chatServer *chat.ChatServer, w http.ResponseWriter, r *http.Request) {
	log.Println("ws conneting....")

	// websocket 기능 추가
	conn, wsErr := upgrader.Upgrade(w, r, nil)
	if wsErr != nil {
		log.Println("wsErr " + wsErr.Error())
		//error 처리 고민...
		return
	}

	var wsReq chat.WsRequest

	if err := conn.ReadJSON(&wsReq); err != nil {
		log.Println("read message error " + err.Error())
		return
	}

	// client의 정보를 가져옴
	user, err := chatHandler.userApp.GetUserByID(wsReq.UserID)
	if err != nil {
		log.Println("get client err " + err.Message)
		conn.WriteJSON(err)
		return
	}

	// client의 정보를 토대로 ChatUser 객체 생성
	chatClient := chat.NewChatUser(user.ID, user.Name, user.Nickname, conn, chatServer)
	// 메시지 보내기를 눌렀을 때는 무조건 새로 생성
	chatRoom := chatServer.CreateRoom(wsReq.ChatRoomName)
	chatClient.ChatRooms[wsReq.ChatRoomName] = chatRoom

	var chatServerReq chat.ChatServerRequest
	chatServerReq.User = chatClient
	chatServerReq.ChatRoomName = wsReq.ChatRoomName

	chatServer.Register <- chatServerReq

	go chatClient.ReadPump()
	go chatClient.WritePump()
}

// GetChatRoomInfo: 채팅룸과 기존메시지(존재하면)를 반환함
func (chatHandler *ChatHandler) GetChatRoomInfo(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	var chatReq chat.ChatRequest

	if err := json.NewDecoder(r.Body).Decode(&chatReq); err != nil {
		restErr := errors.NewBadRequestError("invalid json body " + err.Error())
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
	}
	defer r.Body.Close()

	// host user ID 가져옴
	hostUserID, err := chatHandler.studyPostApp.GetUserIDByPostID(chatReq.StudyPostID)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	// 채팅룸이 기존에 존재하는지 새로 만들어야하는지 확인
	isRoomExist := true
	chatRoom, err := chatHandler.chatApp.GetChatRoom(chatReq.UserID, hostUserID, chatReq.StudyPostID) // chatReq.UserID = clientID
	if err != nil {

		if err.Message == errors.ErrNoRows { // 기존 채팅룸이 존재하지 않으므로 새로운 방 만듬
			chatRoom, restErr := chatHandler.chatApp.SaveChatRoom(chatReq.UserID, hostUserID, chatReq.StudyPostID)
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
			return
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
