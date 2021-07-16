package persistence

import (
	"database/sql"
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/code-wave/go-wave/infrastructure/helpers"
	"github.com/pborman/uuid"
)

type chatRepo struct {
	db *sql.DB
}

func NewChatRepo(db *sql.DB) *chatRepo {
	return &chatRepo{db}
}

// GetChatRoom 채팅룸이 존재하지 않으면 새로 만들기도함
func (c *chatRepo) GetChatRoom(clientID, hostID, studyPostID int64) (*entity.ChatRoom, *errors.RestErr) {
	stmt, err := c.db.Prepare(`
		SELECT *
		FROM chat_room
		WHERE client_id=$1 AND host_id=$2 AND study_post_id=$3;
	`)
	if err != nil {
		return nil, errors.NewInternalServerError("database error " + err.Error())
	}

	var chatRoom entity.ChatRoom

	err = stmt.QueryRow(clientID, hostID, studyPostID).Scan(&chatRoom.ID, &chatRoom.RoomName, &chatRoom.ClientID, &chatRoom.HostID, &chatRoom.StudyPostID)
	if err != nil {
		if err == sql.ErrNoRows { // 채팅룸이 존재하지 않으면 새로 만들고 반환
			noRowsErr := errors.NewNoRowsError()
			return nil, noRowsErr
		}

		return nil, errors.NewInternalServerError("database error " + err.Error())
	}

	return &chatRoom, nil
}

func (c *chatRepo) SaveChatRoom(clientID, hostID, studyPostID int64) (*entity.ChatRoom, *errors.RestErr) {
	stmt, err := c.db.Prepare(`
		INSERT INTO chat_room (room_name, client_id, host_id, study_post_id)
		VALUES ($1, $2, $3, $4)
		RETURNING *;
	`)
	if err != nil {
		return nil, errors.NewInternalServerError("database error " + err.Error())
	}

	var newRoom entity.ChatRoom

	roomName := uuid.New()

	err = stmt.QueryRow(roomName, clientID, hostID, studyPostID).Scan(&newRoom.ID, &newRoom.RoomName, &newRoom.ClientID, &newRoom.HostID, &newRoom.StudyPostID)
	if err != nil {
		return nil, errors.NewInternalServerError("query row error " + err.Error())
	}

	return &newRoom, nil
}

func (c *chatRepo) GetChatRoomByRoomName(roomName string) (*entity.ChatRoom, *errors.RestErr) {
	stmt, err := c.db.Prepare(`
		SELECT *
		FROM chat_room
		WHERE room_name=$1;
	`)
	if err != nil {
		return nil, errors.NewInternalServerError("database error " + err.Error())
	}

	var newRoom entity.ChatRoom

	err = stmt.QueryRow(roomName).Scan(&newRoom.ID, &newRoom.ClientID, &newRoom.RoomName, &newRoom.HostID, &newRoom.StudyPostID)
	if err != nil {
		return nil, errors.NewInternalServerError("query row error " + err.Error())
	}

	return &newRoom, nil
}

func (c *chatRepo) GetChatRoomByID(id int64) (*entity.ChatRoom, *errors.RestErr) {
	stmt, err := c.db.Prepare(`
		SELECT *
		FROM chat_room
		WHERE id=$1;
	`)
	if err != nil {
		return nil, errors.NewInternalServerError("database error " + err.Error())
	}

	var newRoom entity.ChatRoom

	err = stmt.QueryRow(id).Scan(&newRoom.ID, &newRoom.ClientID, &newRoom.RoomName, &newRoom.HostID, &newRoom.StudyPostID)
	if err != nil {
		return nil, errors.NewInternalServerError("query row error " + err.Error())
	}

	return &newRoom, nil
}

func (c *chatRepo) SaveChatMessage(msg *entity.ChatMessage) (*entity.ChatMessage, *errors.RestErr) {
	stmt, err := c.db.Prepare(`
		INSERT INTO chat_message (chat_room_id, chat_room_name, sender_id, sender, message_type, message, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING *;
	`)
	if err != nil {
		return nil, errors.NewInternalServerError("database error " + err.Error())
	}

	now := helpers.GetCurrentTimeForDB()

	newMsg := entity.ChatMessage{}

	err = stmt.QueryRow(msg.ChatRoomID, msg.ChatRoomName, msg.SenderID, msg.Sender, msg.MessageType, msg.Message, now).Scan(&newMsg.ChatRoomID, &newMsg.ChatRoomName, &newMsg.SenderID, &newMsg.Sender, &newMsg.MessageType, &newMsg.Message, &newMsg.CreatedAt)
	if err != nil {
		return nil, errors.NewInternalServerError("queryrow error " + err.Error())
	}

	return &newMsg, nil
}

// TODO: 메시지 query문 수정 필요(user_name 불러오는 query문 작성!)
func (c *chatRepo) GetChatMessages(roomID int64) ([]entity.ChatMessage, *errors.RestErr) {
	stmt, err := c.db.Prepare(`
		SELECT *
		FROM chat_message
		WHERE chat_room_id=$1 ORDER BY created_at DESC;
	`)
	if err != nil {
		return nil, errors.NewInternalServerError("database error " + err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(roomID)
	if err != nil {
		return nil, errors.NewInternalServerError("database error " + err.Error())
	}

	var chatMessages []entity.ChatMessage
	for rows.Next() {
		var chatMessage entity.ChatMessage
		err = rows.Scan(&chatMessage.ChatRoomID, &chatMessage.ChatRoomName, &chatMessage.SenderID, &chatMessage.Sender, &chatMessage.MessageType, &chatMessage.Message, &chatMessage.CreatedAt)
		if err != nil {
			return nil, errors.NewInternalServerError("database error " + err.Error())
		}
		chatMessages = append(chatMessages, chatMessage)
	}

	return chatMessages, nil
}
