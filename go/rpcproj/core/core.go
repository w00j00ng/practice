package core

import (
	"time"
)

type IServer interface {
	Create(args *CreateUser, resp *User) error
	Update(args *UpdateUser, resp *User) error
	Read(userID string, resp *User) error
	Delete(userID string, resp *User) error
	// ReadAll(empty interface{}, resp []*User) error
}

type (
	User struct {
		ID        string
		Name      string
		Age       int64
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	CreateUser struct {
		Name string
		Age  int64
	}

	UpdateUser struct {
		ID   string
		Name string
		Age  int64
	}
)
