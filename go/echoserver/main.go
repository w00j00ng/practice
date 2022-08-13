package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "echoserver/docs"
)

type (
	User struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Age       uint64    `json:"age"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CreateUser struct {
		Name string `json:"name"`
		Age  uint64 `json:"age"`
	}

	UpdateUser struct {
		Name string `json:"name"`
		Age  uint64 `json:"age"`
	}

	UserList struct {
		Users []User `json:"users"`
	}

	Usermap map[string]*User

	Error struct {
		Message string `json:"message"`
	}
)

var (
	users = Usermap{}
)

//----------
// Handlers
//----------

// @Summary     createUser
// @Description create a new user
// @Param       userBody body     CreateUser true "User info body"
// @Success     201      {object} User
// @Failure     400      {object} Error
// @Router      /users [post]
func createUser(c echo.Context) error {
	var (
		reqData     CreateUser
		newID       string    = uuid.New().String()
		currentTime time.Time = time.Now()
		err         error
	)
	if err = c.Bind(&reqData); err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}
	users[newID] = &User{
		ID:        newID,
		Name:      reqData.Name,
		Age:       reqData.Age,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	return c.JSON(http.StatusCreated, users[newID])
}

// @Summary     getUser
// @Description get a user by id
// @Param       id  path     string true "id of the user"
// @Success     200 {object} User
// @Failure     400 {object} Error
// @Router      /users/{id} [get]
func getUser(c echo.Context) error {
	var (
		id  string
		err error
	)
	err = echo.PathParamsBinder(c).String("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, users[id])
}

// @Summary     updateUser
// @Description update user information by id
// @Param       id       path     string     true "id of the user"
// @Param       userBody body     UpdateUser true "User info body"
// @Success     200      {object} User
// @Failure     400      {object} Error
// @Router      /users/{id} [put]
func updateUser(c echo.Context) error {
	var (
		id          string
		reqData     UpdateUser
		bIDExists   bool
		userData    *User
		currentTime time.Time = time.Now()
		err         error
	)
	err = echo.PathParamsBinder(c).String("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}
	if userData, bIDExists = users[id]; !bIDExists {
		return c.JSON(http.StatusBadRequest, Error{Message: "user not found"})
	}
	if err = c.Bind(&reqData); err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}
	if reqData.Name != "" {
		userData.Name = reqData.Name
		userData.UpdatedAt = currentTime
	}
	if reqData.Age != 0 {
		userData.Age = reqData.Age
		userData.UpdatedAt = currentTime
	}
	return c.JSON(http.StatusOK, userData)
}

// @Summary     deleteUser
// @Description delete a user by id
// @Param       id path string true "id of the user"
// @Success     204
// @Failure     400 {object} Error
// @Router      /users/{id} [delete]
func deleteUser(c echo.Context) error {
	var (
		id  string
		err error
	)
	err = echo.PathParamsBinder(c).String("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

// @Summary     getAllUsers
// @Description get all users
// @Success     200 {object} UserList
// @Failure     400 {object} Error
// @Router      /users [get]
func getAllUsers(c echo.Context) error {
	var (
		userList []*User = make([]*User, len(users))
		i        uint64
	)
	for _, user := range users {
		userList[i] = user
		i++
	}
	return c.JSON(http.StatusOK, userList)
}

// @title   My Web API
// @version v1.0.1
// @host    localhost:8000
func main() {
	var (
		e    *echo.Echo
		port int = 8000
	)

	e = echo.New()

	// Routes
	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
