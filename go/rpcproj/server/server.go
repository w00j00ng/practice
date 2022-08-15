package main

import (
	"errors"
	"time"

	"rpcproj/core"

	"github.com/google/uuid"
)

type Usermap map[string]*core.User

var (
	users = Usermap{}
)

type UserManager struct{}

func (m *UserManager) Create(args *core.CreateUser, resp *core.User) error {
	var (
		newID       string    = uuid.New().String()
		currentTime time.Time = time.Now()
	)
	users[newID] = &core.User{
		ID:        newID,
		Name:      args.Name, // (*args).Name
		Age:       args.Age,  // (*args).Age
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	*resp = *users[newID]
	return nil
}

func (m *UserManager) Update(args *core.UpdateUser, resp *core.User) error {
	var (
		userData    *core.User
		bIDExists   bool
		currentTime time.Time = time.Now()
	)
	if userData, bIDExists = users[args.ID]; !bIDExists {
		return errors.New("user not found")
	}
	if args.Name != "" {
		userData.Name = args.Name
		userData.UpdatedAt = currentTime
	}
	if args.Age != 0 {
		userData.Age = args.Age
		userData.UpdatedAt = currentTime
	}
	*resp = *userData
	return nil
}

func (m *UserManager) Read(userID string, resp *core.User) error {
	var (
		userData  *core.User
		bIDExists bool
	)
	if userData, bIDExists = users[userID]; !bIDExists {
		return errors.New("user not found")
	}
	*resp = *userData
	return nil
}

func (m *UserManager) Delete(userID string, message *string) error {
	var (
		bIDExists bool
	)
	if _, bIDExists = users[userID]; !bIDExists {
		return errors.New("user not found")
	}
	delete(users, userID)
	*message = "OK"
	return nil
}

// func (m *UserManager) ReadAll(empty interface{}, resp *[]core.User) error {
// 	*resp = make([]core.User, len(users))
// 	for _, user := range users {
// 		*resp = append(*resp, core.User{
// 			ID:        user.ID,
// 			Name:      user.Name,
// 			Age:       user.Age,
// 			CreatedAt: user.CreatedAt,
// 			UpdatedAt: user.UpdatedAt,
// 		})
// 	}
// 	return nil
// }
