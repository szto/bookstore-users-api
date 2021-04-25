package app

import (
	"github.com/szto/bookstore-users-api/controllers/ping"
	"github.com/szto/bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controlles.SearchUser)
	router.POST("/users", users.CreateUser)
}
