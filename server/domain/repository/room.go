package repository

import (
	"github.com/sekky0905/modern-chat/domain/model"
)

// ChatRoomRepository は、ChatRoom の Repository。
type ChatRoomRepository interface {
	SaveChatRoom(db DB, room *model.ChatRoom) (model.ChatRoomID, error)
	UpdateChatRoom(db DB, id model.ChatRoomID, room *model.ChatRoom) (model.ChatRoomID, error)
	DeleteChatRoom(db DB, id model.ChatRoomID) (model.ChatRoomID, error)
}
