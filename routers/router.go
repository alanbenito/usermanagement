package routers

import (
	"net/http"
	c "usermanagement/controllers"
)

var Routes = map[string] http.HandlerFunc {
	"/users": 			c.UserListHandler,
	"/users/detail":	c.UserDetailHandler,
	"/users/create":	c.UserCreateHandler,
	"/users/update":	c.UserCreateHandler,
	"/users/delete":	c.UserCreateHandler,
}