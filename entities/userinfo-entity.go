package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UID        int `xorm:"pk autoincr"` //tags
	UserName   string
	DepartName string
	CreateAt   *time.Time //time tags
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	if u.CreateAt == nil {
		t := time.Now() //get current time
		u.CreateAt = &t
	}
	return &u
}
