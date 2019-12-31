package datastore

import (
	"github.com/sekky0905/modern-chat/server/domain/model"
	"github.com/sekky0905/modern-chat/server/domain/repository"
)

type commentRepository struct {
}

// NewCommentRepository は、CommentRepository を生成し、返す。
func NewCommentRepositpry() repository.CommentRepository {
	return &commentRepository{}
}

// SaveComment は、Comment を保存する。
func (commentRepository) SaveComment(db repository.DB, comment *model.Comment) (model.CommentID, error) {
	panic("implement me")
}

// UpdateComment は、Comment を更新する。
func (commentRepository) UpdateComment(db repository.DB, id model.CommentID, room *model.Comment) (model.CommentID, error) {
	panic("implement me")
}

// DeleteComment は、Comment を削除する。
func (commentRepository) DeleteComment(db repository.DB, id model.CommentID) (model.CommentID, error) {
	panic("implement me")
}
