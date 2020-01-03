package dto

import "time"

// CommentListDTO は、Comment 用の List の STO.
// TODO 後に無限スクロール系用の項目を追加するかも
type CommentListDTO struct {
	ID         uint
	User       *UserDTO
	ChatRoomID uint
	Content    string
	Liked      []UserDTO
	CreatedAt  time.Time
}
