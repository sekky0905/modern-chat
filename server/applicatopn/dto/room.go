package dto

import "time"

// ChatRoomListDTO は、ChatRoom 用の List の STO.
// TODO 後に無限スクロール系用の項目を追加するかも
type ChatRoomListDTO struct {
	ID        uint
	Title     string
	UserID    uint
	CreatedAt time.Time
}
