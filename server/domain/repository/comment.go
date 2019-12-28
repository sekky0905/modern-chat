package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sekky0905/modern-chat/domain/model"
)

// CommentRepository は、Comment の Repository。
type CommentRepository interface {
	SaveComment(db *gorm.DB, room *model.Comment) (model.CommentID, error)
	UpdateComment(db *gorm.DB, id model.CommentID, room *model.Comment) (model.CommentID, error)
	DeleteComment(db *gorm.DB, id model.CommentID) (model.CommentID, error)
}
