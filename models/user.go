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
	rows, err := dbcon.Db.Query("SELECT username, password FROM tbl_user ")

	if err != nil {
		// error message
		log.Println(err.Error())
	}else{
		for rows.Next(){
			var usr = User{}
			err = rows.Scan(&usr.Username, &usr.Password)
			userList = append(userList,usr)
		}
	}

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
	
	// query insert user
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

func UpdateUser(user User, id int) error {

	// query update user by user id, return error (if any)
	tx, err := dbcon.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	} else {
		defer tx.Rollback()
		query, err := tx.Prepare(s.Join([]string{"UPDATE tbl_user SET ",
			" username = ?",
			" , password = ?", 
			" WHERE id = ?",
		}, " "))

		if err != nil {
			log.Println(err.Error())
			return err
		} else {
			_, err2 := query.Exec(user.Username, user.Password, id)
			if err2 != nil {
				log.Println(err2.Error())
				return err2
			} else {
				err2 = tx.Commit()
			}
		}
		query.Close()
	}

	return err
}

func DeleteUser(id int) error {

	// query delete user by user id, return error (if any)
	tx, err := dbcon.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	} else {
		defer tx.Rollback()
		query, err := tx.Prepare(s.Join([]string{"DELETE FROM tbl_user ",
			" WHERE id = ?",
		}, " "))

		if err != nil {
			log.Println(err.Error())
			return err
		} else {
			_, err2 := query.Exec(id)
			if err2 != nil {
				log.Println(err2.Error())
				return err2
			} else {
				err2 = tx.Commit()
			}
		}
		query.Close()
	}

	return err
}
