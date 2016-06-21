package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	user "usermanagement/models"
)

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	var userList []user.User = user.GetUserList()
	var response user.JSONResponse

	// set response type to application/json
	w.Header().Set("Content-Type", "application/json")

	response.Meta.Status = "success"
	response.Data = userList

	jsonResponse, err := json.Marshal(response)
	if err == nil {
		fmt.Fprintf(w, bytes.NewBuffer(jsonResponse).String())
	}
}

func UserDetailHandler(w http.ResponseWriter, r *http.Request) {
	var response user.JSONResponse
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))

	if err != nil {
		response.Meta.Error = 1
		response.Meta.Status = err.Error()
	}

	u := user.GetUser(userID)

	// set response type to application/json
	w.Header().Set("Content-Type", "application/json")

	response.Meta.Status = "success"
	response.Data = u

	jsonResponse, err := json.Marshal(response)
	fmt.Fprintf(w, bytes.NewBuffer(jsonResponse).String())
}

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var response user.JSONResponse
		var u user.User

		//r.ParseForm()

		u.Username = r.FormValue("username")
		u.Password = r.FormValue("password")

		err := user.CreateUser(u)
		if err != nil {
			response.Meta.Error = 1
			response.Meta.Status = err.Error()
		} else {
			response.Meta.Status = "success"
			response.Data = u
		}

		jsonResponse, err := json.Marshal(response)
		fmt.Fprintf(w, bytes.NewBuffer(jsonResponse).String())
	}
}

func UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		var response user.JSONResponse
		var u user.User
		userID, errs := strconv.Atoi(r.URL.Query().Get("user_id"))
		if errs != nil {
			response.Meta.Error = 1
			response.Meta.Status = errs.Error()
		}

		//r.ParseForm()

		u.Username = r.FormValue("username")
		u.Password = r.FormValue("password")

		err := user.UpdateUser(u, userID)
		if err != nil {
			response.Meta.Error = 1
			response.Meta.Status = err.Error()
		} else {
			response.Meta.Status = "success"
			response.Data = u
		}

		jsonResponse, err := json.Marshal(response)
		fmt.Fprintf(w, bytes.NewBuffer(jsonResponse).String())
	}
}

func UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		var response user.JSONResponse
		userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))

		if err != nil {
			response.Meta.Error = 1
			response.Meta.Status = err.Error()
		}

		err = user.DeleteUser(userID)
		if err != nil {
			response.Meta.Error = 1
			response.Meta.Status = err.Error()
		} else {
			response.Meta.Status = "success"
		}

		jsonResponse, err := json.Marshal(response)
		fmt.Fprintf(w, bytes.NewBuffer(jsonResponse).String())
	}
}
