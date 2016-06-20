package user

import (
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JSONResponse struct {
	Meta struct {
		Error  int    `json:"error"`
		Status string `json:"status"`
	} `json:"meta"`
	Data interface{} `json:"data,omitempty"`	
}

func GetUserList() []User {
	var userList []User = []User {}

	// query get all user, return all users

	return userList
}

func GetUser(id int) User {
	var user User

	// query get user by user id, return request user

	return user
}

func CreateUser(user User) error {
	var err error

	// query update user by user id, return error (if any)

	return err
}

func UpdateUser(user User) error {
	var err error

	// query update user by user id, return error (if any)

	return err
}

func DeleteUser(id int) error {
	var err error

	// query delete user by user id, return error (if any)

	return err
}