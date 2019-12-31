package repository

import (
	"github.com/sekky0905/modern-chat/server/domain/model"
)

// CommentRepository は、Comment の Repository。
type CommentRepository interface {
	SaveComment(db DB, room *model.Comment) (model.CommentID, error)
	UpdateComment(db DB, id model.CommentID, room *model.Comment) (model.CommentID, error)
	DeleteComment(db DB, id model.CommentID) (model.CommentID, error)
}
