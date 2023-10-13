package models

import (
	"github.com/labstack/echo/v4"
)

type CreateUser struct {
	Name     string
	Password string
}

type User struct {
	Name     string
	Password string
	Id       int
}

var users = []User{
	{
		Name:     "Foo",
		Id:       0,
		Password: "abc",
	},
}

type IModelUser interface {
	ReadUserById(id int) (User, *echo.HTTPError)
}

type ModelUser struct{}

func (_ ModelUser) ReadUserById(id int) (User, *echo.HTTPError) {
	for i := range users {
		if users[i].Id == id {
			return users[i], nil
		}
	}

	return User{}, echo.ErrNotFound
}

func (_ ModelUser) CreateUser(createUser CreateUser) (User, *echo.HTTPError) {
	user := User{
		Name:     createUser.Name,
		Password: createUser.Password,
		Id:       len(users),
	}

	users = append(users, user)
	return user, nil
}

func (_ ModelUser) ReadUsers() ([]User, *echo.HTTPError) {
	return users, nil
}
