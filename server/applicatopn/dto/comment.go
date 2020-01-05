package dto

import "time"

// CommentListDTO は、Comment 用の List の DTO
type CommentListDTO struct {
	List []CommentDTOForList
	// TODO 後に無限スクロール系用の項目を追加するかも
}

// CommentDTOForList は、Comment の List 表示に用いる DTO.
type CommentDTOForList struct {
	CommentDTO
	UserDTO
	Liked []LikeDTO
}

// NewCommentDTOForList は、CommentDTOForList を生成し、返す。
func NewCommentDTOForList(commentDTO CommentDTO, userDTO UserDTO, liked []LikeDTO) CommentDTOForList {
	return CommentDTOForList{CommentDTO: commentDTO, UserDTO: userDTO, Liked: liked}
}

// CommentAndUserDTO は、Comment と User の DTO。
type CommentAndUserDTO struct {
	CommentDTO
	UserDTO
}

// CommentDTO は、Comment 用の DTO。
type CommentDTO struct {
	ID         uint
	ChatRoomID uint
	Content    string
	CreatedAt  time.Time
}

// TableName は、構造体に紐づける table の名前を返す。
func (CommentDTO) TableName() string {
	return "comments"
}

// LikeDTO は、Like の DTO。
type LikeDTO struct {
	UserDTO
	LikedAt time.Time
}
