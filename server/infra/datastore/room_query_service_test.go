package datastore

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jinzhu/gorm"
	"github.com/sekky0905/modern-chat/server/applicatopn/dto"
	"github.com/sekky0905/modern-chat/server/applicatopn/query_service"
	"github.com/sekky0905/modern-chat/server/errs"
)

func compareChatRoomDTOForList(x, y dto.ChatRoomDTOForList) bool {
	return x.ChatRoomDTO.ID == y.ChatRoomDTO.ID &&
		x.ChatRoomDTO.Title == y.ChatRoomDTO.Title &&
		x.UserDTO.UserID == y.UserDTO.UserID &&
		x.UserDTO.Name == y.UserDTO.Name
}

func Test_chatRoomQueryService_ListChatRoom(t *testing.T) {
	type args struct {
		db query_service.DB
	}
	tests := []struct {
		name          string
		c             chatRoomQueryService
		args          args
		setupRoomData []ChatRoom
		setupUserData []User
		want          *dto.ChatRoomListDTO
		wantErrFunc   func() error
	}{
		{
			name: "3つのデータが存在する場合、3つのデータを取得すること",
			c:    chatRoomQueryService{},
			args: args{
				db: DBMock.conn,
			},
			setupRoomData: []ChatRoom{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Title:  "test title1",
					UserID: "test user id1",
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					Title:  "test title2",
					UserID: "test user id2",
				},
				{
					Model: gorm.Model{
						ID: 3,
					},
					Title:  "test title3",
					UserID: "test user id3",
				},
			},
			setupUserData: []User{
				{
					ID:   "test user id1",
					Name: "test name1",
				},
				{
					ID:   "test user id2",
					Name: "test name2",
				},
				{
					ID:   "test user id3",
					Name: "test name3",
				},
			},
			want: &dto.ChatRoomListDTO{
				List: []dto.ChatRoomDTOForList{
					{
						ChatRoomDTO: dto.ChatRoomDTO{
							ID:    3,
							Title: "test title3",
						},
						UserDTO: dto.UserDTO{
							UserID: "test user id3",
							Name:   "test name3",
						},
					},
					{
						ChatRoomDTO: dto.ChatRoomDTO{
							ID:    2,
							Title: "test title2",
						},
						UserDTO: dto.UserDTO{
							UserID: "test user id2",
							Name:   "test name2",
						},
					},
					{
						ChatRoomDTO: dto.ChatRoomDTO{
							ID:    1,
							Title: "test title1",
						},
						UserDTO: dto.UserDTO{
							UserID: "test user id1",
							Name:   "test name1",
						},
					},
				},
			},
			wantErrFunc: func() error { return nil },
		},
		{
			name: "1つのデータが存在する場合、1つのデータを取得すること",
			c:    chatRoomQueryService{},
			args: args{
				db: DBMock.conn,
			},
			setupRoomData: []ChatRoom{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Title:  "test title1",
					UserID: "test user id1",
				},
			},
			setupUserData: []User{
				{
					ID:   "test user id1",
					Name: "test name1",
				},
			},
			want: &dto.ChatRoomListDTO{
				List: []dto.ChatRoomDTOForList{
					{
						ChatRoomDTO: dto.ChatRoomDTO{
							ID:    1,
							Title: "test title1",
						},
						UserDTO: dto.UserDTO{
							UserID: "test user id1",
							Name:   "test name1",
						},
					},
				},
			},
			wantErrFunc: func() error { return nil },
		},
		{
			name: "データが存在しない場合、NoSuchDataError を返すこと",
			c:    chatRoomQueryService{},
			args: args{
				db: DBMock.conn,
			},
			setupRoomData: []ChatRoom{},
			setupUserData: []User{},
			want:          nil,
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
			for _, v := range tt.setupRoomData {
				tt.args.db.Create(&v)
				if tt.args.db.NewRecord(v) {
					t.Error("create error")
				}
			}

			for _, v := range tt.setupUserData {
				tt.args.db.Create(&v)
				if tt.args.db.NewRecord(v) {
					t.Error("create error")
				}
			}

			c := chatRoomQueryService{}
			got, err := c.ListChatRoom(tt.args.db)
			wantErr := tt.wantErrFunc()
			if !reflect.DeepEqual(err, wantErr) {
				t.Errorf("chatRoomQueryService.ListChatRoom() error = %v, wantErr %v", err, wantErr)
				return
			}

			opt := cmp.Comparer(compareChatRoomDTOForList)
			if diff := cmp.Diff(tt.want, got, opt); diff != "" {
				t.Errorf("chatRoomQueryService.ListChatRoom() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
