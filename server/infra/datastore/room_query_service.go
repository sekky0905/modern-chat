package datastore

import (
	"github.com/sekky0905/modern-chat/server/applicatopn/dto"
	"github.com/sekky0905/modern-chat/server/applicatopn/query_service"
	"github.com/sekky0905/modern-chat/server/errs"
	"golang.org/x/xerrors"
)

type chatRoomQueryService struct {
}

// NewChatRoomQueryService は、ChatRoomQueryService を生成し、返す。
func NewChatRoomQueryService() query_service.ChatRoomQueryService {
	return &chatRoomQueryService{}
}

// ListChatRoom は、ChatRoom の List を取得する。
func (chatRoomQueryService) ListChatRoom(db query_service.DB) (*dto.ChatRoomListDTO, error) {
	var list []dto.ChatRoomDTOForList
	if err := db.Table("chat_rooms").
		Select("chat_rooms.id, chat_rooms.title, chat_rooms.created_at, users.id, users.name").
		Joins("INNER JOIN users ON chat_rooms.user_id = users.id").Order("chat_rooms.id desc"). // TODO paging 実装時に Limit 付与すること
		Scan(&list).
		Error; err != nil {
		return nil, xerrors.Errorf("failed to find chat rooms: %w", err)
	}

	if len(list) == 0 {
		detail := make(map[string]string)
		detail["list"] = "empty"
		err := errs.NewNoSuchDataError(detail, nil)
		return nil, err
	}

	return &dto.ChatRoomListDTO{
		List: list,
	}, nil
}
