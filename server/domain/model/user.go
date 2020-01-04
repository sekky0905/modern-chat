package model

// UserID は、User の ID。
type UserID string

// String は、UserID を string 型にして返す。
func (id UserID) String() string {
	return string(id)
}

// User は、ユーザーを表す。
type User struct {
	ID   UserID
	Name string
}
