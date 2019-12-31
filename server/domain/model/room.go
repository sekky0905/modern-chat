package model

import (
	"time"

	"github.com/sekky0905/modern-chat/server/util"
)

// ChatRoomID は、ChatRoom のID。
type ChatRoomID uint32

// String は、ChatRoomID を string 型にして返す。
func (id ChatRoomID) String() string {
	return string(id)
}

// ChatRoom は、チャットルームを表す。
type ChatRoom struct {
	ID        ChatRoomID
	Title     string
	UserID    UserID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewChatRoom は、ChatRoom を生成し、返す。
func NewChatRoom(title string, userID UserID) *ChatRoom {
	return &ChatRoom{
		Title:     title,
		UserID:    userID,
		CreatedAt: util.Now(),
		UpdatedAt: util.Now(),
	}
}
