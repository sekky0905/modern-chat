package dto

// UserDTO は、User 用の DTO。
type UserDTO struct {
	UserID string `gorm:"column:id"`
	Name   string
}

// TableName は、構造体に紐づける table の名前を返す。
func (UserDTO) TableName() string {
	return "users"
}
