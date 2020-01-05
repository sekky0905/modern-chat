package datastore

import (
	"github.com/sekky0905/modern-chat/server/domain/model"
	"github.com/sekky0905/modern-chat/server/domain/repository"
	"golang.org/x/xerrors"
)

type commentRepository struct {
}

// NewCommentRepository は、CommentRepository を生成し、返す。
func NewCommentRepositpry() repository.CommentRepository {
	return &commentRepository{}
}

// SaveComment は、Comment を保存する。
func (commentRepository) SaveComment(db repository.DB, comment *model.Comment) (model.CommentID, error) {
	commentDTO := newCommentTranslateFromDomainModel(comment)

	if err := db.Create(&commentDTO).Error; err != nil {
		return 0, xerrors.New("failed to create comment")
	}

	if db.NewRecord(commentDTO) {
		return 0, xerrors.New("failed to create comment")
	}

	likeDTO := newLikesTranslateFromDomainModel(commentDTO.ID, comment.Liked)
	for _, dto := range likeDTO {
		db.Create(&dto)
		if db.NewRecord(dto) {
			return 0, xerrors.New("failed to create like")
		}
	}

	return newCommentIDFromUint(commentDTO.ID), nil
}

// DeleteComment は、Comment を削除する。
func (commentRepository) DeleteComment(db repository.DB, comment *model.Comment) model.CommentID {
	commentDTO := newCommentTranslateFromDomainModel(comment)
	db.Unscoped().Delete(&commentDTO)
	likeDTO := newLikesTranslateFromDomainModel(commentDTO.ID, comment.Liked)
	db.Unscoped().Delete(&likeDTO)
	return comment.ID
}
