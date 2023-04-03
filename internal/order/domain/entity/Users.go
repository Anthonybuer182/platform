package entity

import (
	"time"
)

type Users struct {
	UserId   int32
	UserName string // shadow field
	Passwd   string
	Email    string
	CreateOn time.Time
}

func NewUsers(userId int32, userName string, passwd string, email string, createOn time.Time) *Users {
	return &Users{
		UserId:   userId,
		UserName: userName,
		Passwd:   passwd,
		Email:    email,
		CreateOn: createOn,
	}
}
