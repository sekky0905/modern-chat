package datastore

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jinzhu/gorm"
	"github.com/sekky0905/modern-chat/server/domain/model"
	"github.com/sekky0905/modern-chat/server/domain/repository"
)

func compareChatRoom(x, y *ChatRoom) bool {
	if x == nil && y == nil {
		return true
	}

	return x != nil && y != nil &&
		x.Model.ID == y.Model.ID &&
		x.UserID == y.UserID
}

func Test_chatRoomRepository_SaveChatRoom(t *testing.T) {
	DBMock.truncateTables()

	type args struct {
		db   repository.DB
		room *model.ChatRoom
	}
	tests := []struct {
		name     string
		c        chatRoomRepository
		args     args
		wantID   model.ChatRoomID
		want     *ChatRoom
		wantFunc func(target *ChatRoom) bool
		wantErr  bool
	}{
		{
			name: "適切なデータを与えると、データが適切に格納されること",
			c:    chatRoomRepository{},
			args: args{
				db: DBMock.conn,
				room: &model.ChatRoom{
					Title:  "test title",
					UserID: "test user id",
				},
			},
			want: &ChatRoom{
				Model: gorm.Model{
					ID: 1,
				},
				Title:  "test title",
				UserID: "test user id",
			},
			wantFunc: func(target *ChatRoom) bool {
				return !target.CreatedAt.IsZero() && !target.UpdatedAt.IsZero() && target.DeletedAt == nil
			},
			wantID:  model.ChatRoomID(1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := chatRoomRepository{}
			gotID, err := c.SaveChatRoom(tt.args.db, tt.args.room)
			if (err != nil) != tt.wantErr {
				t.Errorf("chatRoomRepository.SaveChatRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantID != gotID {
				t.Errorf("chatRoomRepository.SaveChatRoom() error = %v, wantErr %v", err, tt.wantErr)
			}

			var got ChatRoom
			tt.args.db.Last(&got)

			opt := cmp.Comparer(compareChatRoom)
			if diff := cmp.Diff(tt.want, &got, opt); diff != "" && tt.wantFunc(&got) {
				t.Errorf("chatRoomRepository.SaveChatRoom() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
