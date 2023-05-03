package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []user{
	{ID: "1", Username: "ajesk", Password: "abcd"},
	{ID: "2", Username: "notajesk", Password: "abcd"},
	{ID: "3", Username: "third", Password: "abcd"},
}

func getUsers(c *gin.Context) {
	usrs := getDbUsers()
	c.IndentedJSON(http.StatusOK, usrs)
}

func getUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		panic(err)
	}

	user := getDbUserById(id)
	c.IndentedJSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	insertUser(newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
