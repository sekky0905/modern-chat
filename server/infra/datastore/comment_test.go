package datastore

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jinzhu/gorm"
	"github.com/sekky0905/modern-chat/server/domain/model"
	"github.com/sekky0905/modern-chat/server/domain/repository"
)

func compareComment(x, y *Comment) bool {
	if x == nil && y == nil {
		return true
	}

	return x != nil && y != nil &&
		x.Model.ID == y.Model.ID &&
		x.UserID == y.UserID &&
		x.ChatRoomID == y.ChatRoomID &&
		x.Content == y.Content
}

func compareLike(x, y Like) bool {
	return x.UserID == y.UserID &&
		x.CommentID == y.CommentID
}

func Test_commentRepository_SaveComment(t *testing.T) {
	DBMock.truncateTables()

	type args struct {
		db      repository.DB
		comment *model.Comment
	}
	tests := []struct {
		name             string
		c                commentRepository
		args             args
		wantID           model.CommentID
		wantComment      *Comment
		checkCommentFunc func(target *Comment) bool
		wantLikes        []Like
		checkLikesFunc   func(target []Like) bool
		wantErr          bool
	}{
		{
			name: "適切なデータを与えると、データが適切に格納されること",
			c:    commentRepository{},
			args: args{
				db: DBMock.conn,
				comment: &model.Comment{
					UserID:     "test user id1",
					ChatRoomID: model.ChatRoomID(1),
					Content:    "test content",
					Liked:      []model.UserID{"test user id2", "test user id3", "test user id4"},
				},
			},
			wantID: model.CommentID(1),
			wantComment: &Comment{
				Model: gorm.Model{
					ID: 1,
				},
				UserID:     "test user id1",
				ChatRoomID: 1,
				Content:    "test content",
			},
			checkCommentFunc: func(target *Comment) bool {
				return !target.CreatedAt.IsZero() && !target.UpdatedAt.IsZero() && target.DeletedAt == nil
			},
			wantLikes: []Like{
				{
					UserID:    "test user id2",
					CommentID: 1,
				},
				{
					UserID:    "test user id3",
					CommentID: 1,
				},
				{
					UserID:    "test user id4",
					CommentID: 1,
				},
			},
			checkLikesFunc: func(target []Like) bool {
				for _, like := range target {
					if like.CreatedAt.IsZero() || like.UpdatedAt.IsZero() || like.DeletedAt != nil {
						return false
					}
				}
				return true
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := commentRepository{}
			gotID, err := c.SaveComment(tt.args.db, tt.args.comment)
			if (err != nil) != tt.wantErr {
				t.Errorf("commentRepository.SaveComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotID != tt.wantID {
				t.Errorf("commentRepository.SaveComment() = %v, want %v", gotID, tt.wantID)
			}

			var gotComment Comment
			tt.args.db.Last(&gotComment)

			optForComment := cmp.Comparer(compareComment)
			if diff := cmp.Diff(tt.wantComment, &gotComment, optForComment); diff != "" && tt.checkCommentFunc(&gotComment) {
				t.Errorf("commentRepository.SaveComment() mismatch (-want +got):\n%s", diff)
			}

			var gotLikes []Like
			tt.args.db.Order("user_id").Find(&gotLikes)

			optForLike := cmp.Comparer(compareLike)
			if diff := cmp.Diff(tt.wantLikes, gotLikes, optForLike); diff != "" && tt.checkLikesFunc(gotLikes) {
				t.Errorf("commentRepository.SaveComment() mismatch (+got -want) :\n%s", diff)
			}
		})
	}
}
