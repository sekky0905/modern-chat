package model

import "time"

// ChatRoomID は、ChatRoom のID。
type ChatRoomID uint32

// String は、UserID を string 型にして返す。
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
