package dto

import "time"

// ChatRoomListDTO は、ChatRoom 用の List の DTO
type ChatRoomListDTO struct {
	List []ChatRoomDTOForList
	// TODO 後に無限スクロール系用の項目を追加するかも
}

// ChatRoomDTOForList は、ChatRoom の List 表示に用いる DTO.
type ChatRoomDTOForList struct {
	ChatRoomDTO
	UserDTO
}

// ChatRoomDTO は、ChatRoom 用の DTO。
type ChatRoomDTO struct {
	ID        uint
	Title     string
	CreatedAt time.Time
}

// TableName は、構造体に紐づける table の名前を返す。
func (ChatRoomDTO) TableName() string {
	return "chat_rooms"
}
