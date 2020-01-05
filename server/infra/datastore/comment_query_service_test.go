package datastore

import (
	"reflect"
	"testing"

	"github.com/sekky0905/modern-chat/server/errs"

	"github.com/google/go-cmp/cmp"
	"github.com/jinzhu/gorm"
	"github.com/sekky0905/modern-chat/server/applicatopn/dto"
	"github.com/sekky0905/modern-chat/server/applicatopn/query_service"
	"golang.org/x/xerrors"
)

func compareCommentDTO(x, y dto.CommentDTO) bool {
	return x.ID == y.ID &&
		x.ChatRoomID == y.ChatRoomID &&
		x.Content == y.Content
}

func compareUserDTO(x, y dto.UserDTO) bool {
	return x.UserID == y.UserID &&
		x.Name == y.Name
}

func compareLikeDTO(x, y dto.LikeDTO) bool {
	return compareUserDTO(x.UserDTO, y.UserDTO)

}

func Test_commentQueryService_ListComment(t *testing.T) {
	type args struct {
		db         query_service.DB
		chatRoomID uint
	}
	tests := []struct {
		name             string
		q                commentQueryService
		args             args
		setupCommentData []Comment
		setupUserData    []User
		setupLikeData    []Like
		want             *dto.CommentListDTO
		wantErrFunc      func() error
	}{
		{
			name: "3つのデータが存在する場合、3つのデータを取得すること",
			q:    commentQueryService{},
			args: args{
				db:         DBMock.conn,
				chatRoomID: 1,
			},
			setupCommentData: []Comment{
				{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:     "test user id1",
					ChatRoomID: 1,
					Content:    "test content1",
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					UserID:     "test user id2",
					ChatRoomID: 1,
					Content:    "test content2",
				},
				{
					Model: gorm.Model{
						ID: 3,
					},
					UserID:     "test user id3",
					ChatRoomID: 1,
					Content:    "test content3",
				},
			},
			setupUserData: []User{
				// Comment した側用
				{
					ID:   "test user id1",
					Name: "test user name1",
				},
				{
					ID:   "test user id2",
					Name: "test user name2",
				},
				{
					ID:   "test user id3",
					Name: "test user name3",
				},
				// Like した側用
				// Comment1 への Like用
				{
					ID:   "test user id4",
					Name: "test user name4",
				},
				{
					ID:   "test user id5",
					Name: "test user name5",
				},
				// Comment2 への Like用
				{
					ID:   "test user id6",
					Name: "test user name6",
				},
				{
					ID:   "test user id7",
					Name: "test user name7",
				},
				// Comment3への Like用
				{
					ID:   "test user id8",
					Name: "test user name8",
				},
				{
					ID:   "test user id9",
					Name: "test user name9",
				},
			},
			setupLikeData: []Like{
				{
					CommentID: 1,
					UserID:    "test user id4",
				},
				{
					CommentID: 1,
					UserID:    "test user id5",
				},
				{
					CommentID: 2,
					UserID:    "test user id6",
				},
				{
					CommentID: 2,
					UserID:    "test user id7",
				},
				{
					CommentID: 3,
					UserID:    "test user id8",
				},
				{
					CommentID: 3,
					UserID:    "test user id9",
				},
			},
			want: &dto.CommentListDTO{
				List: []dto.CommentDTOForList{
					{
						CommentDTO: dto.CommentDTO{
							ID:         3,
							ChatRoomID: 1,
							Content:    "test content3",
						},
						UserDTO: dto.UserDTO{
							UserID: "test user id3",
							Name:   "test user name3",
						},
						Liked: []dto.LikeDTO{
							{
								UserDTO: dto.UserDTO{
									UserID: "test user id8",
									Name:   "test user name8",
								},
							},
							{
								UserDTO: dto.UserDTO{
									UserID: "test user id9",
									Name:   "test user name9",
								},
							},
						},
					},
					{
						CommentDTO: dto.CommentDTO{
							ID:         2,
							ChatRoomID: 1,
							Content:    "test content2",
						},
						UserDTO: dto.UserDTO{
							UserID: "test user id2",
							Name:   "test user name2",
						},
						Liked: []dto.LikeDTO{
							{
								UserDTO: dto.UserDTO{
									UserID: "test user id6",
									Name:   "test user name6",
								},
							},
							{
								UserDTO: dto.UserDTO{
									UserID: "test user id7",
									Name:   "test user name7",
								},
							},
						},
					},
					{
						CommentDTO: dto.CommentDTO{
							ID:         1,
							ChatRoomID: 1,
							Content:    "test content1",
						},
						UserDTO: dto.UserDTO{
							UserID: "test user id1",
							Name:   "test user name1",
						},
						Liked: []dto.LikeDTO{
							{
								UserDTO: dto.UserDTO{
									UserID: "test user id4",
									Name:   "test user name4",
								},
							},
							{
								UserDTO: dto.UserDTO{
									UserID: "test user id5",
									Name:   "test user name5",
								},
							},
						},
					},
				},
			},
			wantErrFunc: func() error { return nil },
		},
		{
			name: "1つのデータが存在する場合、1つのデータを取得すること",
			q:    commentQueryService{},
			args: args{
				db:         DBMock.conn,
				chatRoomID: 1,
			},
			setupCommentData: []Comment{
				{
					Model: gorm.Model{
						ID: 1,
					},
					UserID:     "test user id1",
					ChatRoomID: 1,
					Content:    "test content1",
				},
			},
			setupUserData: []User{
				// Comment した側用
				{
					ID:   "test user id1",
					Name: "test user name1",
				},
				// Like した側用
				// Comment1 への Like用
				{
					ID:   "test user id2",
					Name: "test user name2",
				},
			},
			setupLikeData: []Like{
				{
					CommentID: 1,
					UserID:    "test user id2",
				},
			},
			want: &dto.CommentListDTO{
				List: []dto.CommentDTOForList{
					{
						CommentDTO: dto.CommentDTO{
							ID:         1,
							ChatRoomID: 1,
							Content:    "test content1",
						},
						UserDTO: dto.UserDTO{
							UserID: "test user id1",
							Name:   "test user name1",
						},
						Liked: []dto.LikeDTO{
							{
								UserDTO: dto.UserDTO{
									UserID: "test user id2",
									Name:   "test user name2",
								},
							},
						},
					},
				},
			},
			wantErrFunc: func() error { return nil },
		},
		{
			name: "データが存在しない場合、NoSuchDataError を返すこと",
			q:    commentQueryService{},
			args: args{
				db:         DBMock.conn,
				chatRoomID: 1,
			},
			want: nil,
			wantErrFunc: func() error {
				detail := make(map[string]string)
				detail["list"] = "empty"
				return errs.NewNoSuchDataError(detail, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DBMock.truncateTables()

			// set up
			for _, v := range tt.setupCommentData {
				if err := tt.args.db.Create(&v).Error; err != nil {
					t.Error("create error")
				}
				if tt.args.db.NewRecord(v) {
					t.Error("create error")
				}
			}

			for _, v := range tt.setupUserData {
				if err := tt.args.db.Create(&v).Error; err != nil {
					t.Error("create error")
				}
				if tt.args.db.NewRecord(v) {
					t.Error("create error")
				}
			}

			for _, v := range tt.setupLikeData {
				if err := tt.args.db.Create(&v).Error; err != nil {
					t.Error("create error")
				}
				if tt.args.db.NewRecord(v) {
					t.Error("create error")
				}
			}

			q := commentQueryService{}
			got, err := q.ListCommentByChatRoomID(tt.args.db, tt.args.chatRoomID)
			wantErr := tt.wantErrFunc()
			if !reflect.DeepEqual(xerrors.Unwrap(err), wantErr) {
				t.Errorf("commentQueryService.ListComment() error = %v, wantErr %+v", xerrors.Unwrap(err), wantErr)
				return
			}

			if tt.want == nil {
				return
			}

			for i, gotElm := range got.List {
				opt := cmp.Comparer(compareCommentDTO)
				if diff := cmp.Diff(tt.want.List[i].CommentDTO, gotElm.CommentDTO, opt); diff != "" {
					t.Errorf("commentQueryService.ListComment() mismatch (-want +got):\n%s", diff)
				}

				opt = cmp.Comparer(compareUserDTO)
				if diff := cmp.Diff(tt.want.List[i].UserDTO, gotElm.UserDTO, opt); diff != "" {
					t.Errorf("commentQueryService.ListComment() mismatch (-want +got):\n%s", diff)
				}

				opt = cmp.Comparer(compareLikeDTO)
				if diff := cmp.Diff(tt.want.List[i].Liked, gotElm.Liked, opt); diff != "" {
					t.Errorf("commentQueryService.ListComment() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
