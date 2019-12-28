package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sekky0905/modern-chat/domain/model"
)

// ChatRoomRepository は、ChatRoom の Repository。
type ChatRoomRepository interface {
	SaveChatRoom(db *gorm.DB, room *model.ChatRoom) (model.ChatRoomID, error)
	UpdateChatRoom(db *gorm.DB, id model.ChatRoomID, room *model.ChatRoom) (model.ChatRoomID, error)
	DeleteChatRoom(db *gorm.DB, id model.ChatRoomID) (model.ChatRoomID, error)
}
