package entity

import "github.com/google/uuid"

type Users struct {
	ID       uuid.UUID
	UserName string // shadow field
	Passwd   string
	Email    string
}

func NewUsers(userName string, passwd string, email string) *Users {
	return &Users{
		ID:       uuid.New(),
		UserName: userName,
		Passwd:   passwd,
		Email:    email,
	}
}
