package datastore

import (
	"github.com/sekky0905/modern-chat/server/domain/model"
	"github.com/sekky0905/modern-chat/server/domain/repository"
	"golang.org/x/xerrors"
)

type chatRoomRepository struct {
}

// NewChatRoomRepositpry は、ChatRoomRepository を生成し、返す。
func NewChatRoomRepositpry() repository.ChatRoomRepository {
	return &chatRoomRepository{}
}

// SaveChatRoom は、ChatRoom を保存する。
func (chatRoomRepository) SaveChatRoom(db repository.DB, room *model.ChatRoom) (model.ChatRoomID, error) {
	dto := newChatRoomTranslateFromDomainModel(room)

	db.Create(&dto)

	if db.NewRecord(dto) {
		return 0, xerrors.New("failed to create chat room")
	}

	return newChatRoomIDFromUint(dto.ID), nil
}
