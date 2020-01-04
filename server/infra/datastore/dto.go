package datastore

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sekky0905/modern-chat/server/domain/model"
)

type AtFields struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// User は、ユーザーを表す。
type User struct {
	ID   uint32 `gorm:"type:varchar(36)"`
	Name string
	AtFields
}

// TranslateToDomainModel は、domain model に移し替える。
func (u *User) TranslateToDomainModel() *model.User {
	return &model.User{
		ID:   model.UserID(u.ID),
		Name: u.Name,
	}
}

// ChatRoom は、チャットルームを表す。
type ChatRoom struct {
	gorm.Model
	Title  string `gorm:"type:varchar(20)"`
	UserID string `gorm:"type:varchar(36)"`
}

// TranslateToDomainModel は、domain model に移し替える。
func (c *ChatRoom) TranslateToDomainModel() *model.ChatRoom {
	return &model.ChatRoom{
		ID:        model.ChatRoomID(c.ID),
		Title:     c.Title,
		UserID:    model.UserID(c.UserID),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

// newChatRoomTranslateFromDomainModel は、Domain Model から ChatRoom を生成し、返す。
func newChatRoomTranslateFromDomainModel(in *model.ChatRoom) ChatRoom {
	return ChatRoom{
		Title:  in.Title,
		UserID: in.UserID.String(),
	}
}

// newChatRoomIDFromUint は、Uint の値から ChatRoomID を生成し、返す。
func newChatRoomIDFromUint(id uint) model.ChatRoomID {
	return model.ChatRoomID(id)
}

// Comment は、コメントを表す。
type Comment struct {
	gorm.Model
	UserID     string `gorm:"type:varchar(36)"`
	ChatRoomID uint   `gorm:"INT UNSIGNED NOT NULL"`
	Content    string `gorm:"type:varchar(200)"`
}

// TranslateToDomainModel は、domain model に移し替える。
func (c *Comment) TranslateToDomainModel(likes []Like) *model.Comment {
	n := len(likes)
	s := make([]model.UserID, n, n)

	for i, v := range likes {
		s[i] = model.UserID(v.UserID)
	}

	return &model.Comment{
		ID:         model.CommentID(c.ID),
		UserID:     model.UserID(c.UserID),
		ChatRoomID: model.ChatRoomID(c.ChatRoomID),
		Liked:      s,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

// newCommentTranslateFromDomainModel は、Domain Model から Comment を生成し、返す。
func newCommentTranslateFromDomainModel(in *model.Comment) Comment {
	return Comment{
		Model: gorm.Model{
			ID:        in.ID.Uint(),
			CreatedAt: in.CreatedAt,
			UpdatedAt: in.UpdatedAt,
		},
		UserID:     in.UserID.String(),
		ChatRoomID: in.ChatRoomID.Uint(),
		Content:    in.Content,
	}
}

// newCommentIDFromUint は、Uint の値から CommentID を生成し、返す。
func newCommentIDFromUint(id uint) model.CommentID {
	return model.CommentID(id)
}

// Like は、いいねを表す。
type Like struct {
	UserID    string `gorm:"type:varchar(36);primary_key;auto_increment:false"`
	CommentID uint   `gorm:"NOT NULL;default:0;primary_key;auto_increment:false"`
	AtFields
}

// newLikesTranslateFromDomainModel は、Domain Model から Like を生成し、返す。
func newLikesTranslateFromDomainModel(commentID uint, userIDs []model.UserID) []Like {
	n := len(userIDs)
	likes := make([]Like, n, n)
	for i, userID := range userIDs {
		like := Like{
			UserID:    userID.String(),
			CommentID: commentID,
		}
		likes[i] = like
	}

	return likes
}
