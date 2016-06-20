package models

import (
	"log"
	s "strings"
	dbcon "usermanagement/db"
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
	var userList []User = []User{}

	// query get all user, return all users

	return userList
}

func GetUser(id int) User {
	var user User

	// query get user by user id, return request user
	err := dbcon.Db.QueryRow("SELECT username, password FROM tbl_user WHERE id = ?", id).Scan(&user.Username, &user.Password)

	if err != nil {
		// error message
		log.Println(err.Error())
	}

	return user
}

func CreateUser(user User) error {
	tx, err := dbcon.Db.Begin()
	if err != nil {
		log.Println("error koneksi ", err.Error())
		return err
	} else {
		defer tx.Rollback()
		query, err := tx.Prepare(s.Join([]string{"INSERT INTO tbl_user ",
			"(username, password) VALUES(?, ?) "}, " "))

		if err != nil {
			log.Println("error query ", err.Error())
			return err
		} else {
			_, err2 := query.Exec(user.Username, user.Password)
			if err2 != nil {
				log.Println("error param ", err2.Error())
				return err2
			} else {
				err2 = tx.Commit()
			}
		}
		query.Close()
	}

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
