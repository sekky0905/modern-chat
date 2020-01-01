package repository

import (
	"github.com/sekky0905/modern-chat/server/domain/model"
)

// CommentRepository は、Comment の Repository。
type CommentRepository interface {
	SaveComment(db DB, room *model.Comment) (model.CommentID, error)
	DeleteComment(db DB, comment *model.Comment) model.CommentID
}
