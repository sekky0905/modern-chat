package query_service

import "github.com/sekky0905/modern-chat/server/applicatopn/dto"

// CommentQueryService は、Comment 用の QueryService。
type CommentQueryService interface {
	ListCommentByChatRoomID(db DB, chatRoomID uint) (*dto.CommentListDTO, error)
}
