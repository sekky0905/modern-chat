package datastore

import (
	"github.com/sekky0905/modern-chat/server/domain/model"
	"github.com/sekky0905/modern-chat/server/domain/repository"
	"golang.org/x/xerrors"
)

type chatRoomRepository struct {
}

// NewChatRoomRepository は、ChatRoomRepository を生成し、返す。
func NewChatRoomRepositpry() repository.ChatRoomRepository {
	return &chatRoomRepository{}
}

// SaveChatRoom は、ChatRoom を保存する。
func (chatRoomRepository) SaveChatRoom(db repository.DB, room *model.ChatRoom) (*model.ChatRoom, error) {
	dto := newChatRoomTranslateFromDomainModel(room)

	if err := db.Create(&dto).Error; err != nil {
		return nil, xerrors.New("failed to create chat room")
	}

	if db.NewRecord(dto) {
		return nil, xerrors.New("failed to create chat room")
	}

	return newChatRoomDomainModelFromDTO(dto), nil
}
