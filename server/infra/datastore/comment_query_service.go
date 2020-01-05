package datastore

import (
	"fmt"

	"github.com/sekky0905/modern-chat/server/applicatopn/dto"
	"github.com/sekky0905/modern-chat/server/applicatopn/query_service"
	"github.com/sekky0905/modern-chat/server/errs"
	"github.com/sekky0905/modern-chat/server/util"
	"golang.org/x/xerrors"
)

type commentQueryService struct {
}

// NewCommentQueryService は、CommentQueryService を生成し、返す。
func NewCommentQueryService() query_service.CommentQueryService {
	return &commentQueryService{}
}

// ListCommentByChatRoomID は、Comment の List を取得する。
func (q commentQueryService) ListCommentByChatRoomID(db query_service.DB, chatRoomID uint) (*dto.CommentListDTO, error) {
	comments, err := q.listCommentAndUserDTO(db, chatRoomID)
	if err != nil {
		return nil, xerrors.Errorf("failed to list comment and user dto: %w", err)
	}

	n := len(comments)
	commentList := make([]dto.CommentDTOForList, n, n)

	for i, comment := range comments {
		likes, err := q.listLikedUser(db, comment.ID)
		if err != nil {
			// Like が取得できなかっただけなので、ログに残すに止める
			phrase := fmt.Sprintf("failed to list likes. err = %+v", err)
			util.Logger().Error(phrase)
		}

		commentDTO := dto.NewCommentDTOForList(comment.CommentDTO, comment.UserDTO, likes)
		commentList[i] = commentDTO
	}

	return &dto.CommentListDTO{
		List: commentList,
	}, nil
}

func (commentQueryService) listCommentAndUserDTO(db query_service.DB, chatRoomID uint) ([]dto.CommentAndUserDTO, error) {
	var list []dto.CommentAndUserDTO
	if err := db.Table("comments").
		Select("comments.id, comments.chat_room_id, comments.content, comments.created_at, users.id, users.name").
		Joins("INNER JOIN users ON comments.user_id = users.id").
		//Where("comments.chat_room_id = ?", chatRoomID).
		Order("comments.id desc"). // TODO paging 実装時に Limit 付与すること
		Scan(&list).
		Error; err != nil {
		return nil, xerrors.Errorf("failed to find comment and user dto: %w", err)
	}

	if len(list) == 0 {
		detail := make(map[string]string)
		detail["list"] = "empty"
		err := errs.NewNoSuchDataError(detail, nil)
		return nil, err
	}

	return list, nil
}

func (commentQueryService) listLikedUser(db query_service.DB, commentID uint) ([]dto.LikeDTO, error) {
	var list []dto.LikeDTO
	if err := db.Table("likes").
		Select("likes.comment_id, likes.user_id, likes.created_at, users.id, users.name").
		Joins("INNER JOIN users ON likes.user_id = users.id").
		Where("likes.comment_id = ?", commentID).
		Order("likes.created_at desc"). // TODO paging 実装時に Limit 付与すること
		Scan(&list).
		Error; err != nil {
		return nil, xerrors.Errorf("failed to find like dto: %w", err)
	}

	if len(list) == 0 {
		detail := make(map[string]string)
		detail["list"] = "empty"
		err := errs.NewNoSuchDataError(detail, nil)
		return nil, err
	}

	return list, nil
}
