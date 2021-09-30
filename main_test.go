package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGETUsers(t *testing.T) {
	r := gin.Default()
	r.GET("/getUsers", getUsers)
}

/////

func TestDeleteUser(t *testing.T) {
	r := gin.Default()
	r.DELETE("/delete/John", deleteUser)
}

func TestAddUser(t *testing.T) {
	r := gin.Default()
	r.POST("/create", createUser)
}

func TestUpdateUser(t *testing.T) {
	r := gin.Default()
	r.POST("/update", updateUser)
}
