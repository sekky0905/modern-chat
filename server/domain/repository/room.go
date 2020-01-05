package repository

import (
	"github.com/sekky0905/modern-chat/server/domain/model"
)

// ChatRoomRepository は、ChatRoom の Repository。
type ChatRoomRepository interface {
	SaveChatRoom(db DB, room *model.ChatRoom) (*model.ChatRoom, error)
}
