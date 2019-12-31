package model

import (
	"time"

	"github.com/sekky0905/modern-chat/server/util"
)

// CommentID は、Comment の ID。
type CommentID uint32

// Comment は、コメントを表す。
type Comment struct {
	ID         CommentID
	UserID     UserID
	ChatRoomID ChatRoomID
	Content    string
	Liked      []UserID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// NewComment は、Comment を生成し、返す。
func NewComment(userID UserID, chatRoomID ChatRoomID, content string) *Comment {
	return &Comment{
		UserID:     userID,
		ChatRoomID: chatRoomID,
		Content:    content,
		Liked:      []UserID{},
		CreatedAt:  util.Now(),
		UpdatedAt:  util.Now(),
	}
}
