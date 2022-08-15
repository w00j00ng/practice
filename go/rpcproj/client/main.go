package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpcproj/core"
)

func main() {
	var (
		portNumber int64 = 7000
		rpcClient  *rpc.Client
		createUser core.CreateUser
		updateUser core.UpdateUser
		user       core.User
		// allUsers   []core.User
		userID  string
		message string
		err     error
	)

	log.SetFlags(log.Lshortfile)

	rpcClient, err = rpc.DialHTTP("tcp", fmt.Sprintf(":%d", portNumber))
	if err != nil {
		log.Fatalln(err.Error())
	}

	createUser = core.CreateUser{
		Name: "James",
		Age:  25,
	}
	err = rpcClient.Call("UserManager.Create", &createUser, &user)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("USER CREATED: %#v\n", user)

	userID = user.ID

	err = rpcClient.Call("UserManager.Read", userID, &user)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("READ USER: %#v\n", user)

	updateUser = core.UpdateUser{
		ID:  userID,
		Age: 26,
	}

	err = rpcClient.Call("UserManager.Update", &updateUser, &user)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("USER UPDATED: %#v\n", user)

	err = rpcClient.Call("UserManager.Read", userID, &user)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("READ USER: %#v\n", user)

	createUser = core.CreateUser{
		Name: "Jessica",
		Age:  27,
	}
	err = rpcClient.Call("UserManager.Create", &createUser, &user)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("USER CREATED: %#v\n", user)

	// rpcClient.Call("UserManager.ReadAll", nil, &allUsers)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// log.Printf("GET ALL USERS: %#v\n", &allUsers)

	err = rpcClient.Call("UserManager.Delete", userID, &message)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("USER DELETED: %s\n", message)

	// rpcClient.Call("UserManager.ReadAll", nil, &allUsers)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// log.Printf("GET ALL USERS: %#v\n", &allUsers)
}
