package query_service

import "github.com/sekky0905/modern-chat/server/applicatopn/dto"

// ChatRoomQueryService は、ChatRoom 用の QueryService。
type ChatRoomQueryService interface {
	ListChatRoom(db DB) (*dto.ChatRoomListDTO, error)
}
