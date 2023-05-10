package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	usrs := getDbUsers()
	c.IndentedJSON(http.StatusOK, usrs)
}

func getUserById(c *gin.Context) {
	idStr := c.Param("userId")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := getDbUserById(id)
	c.IndentedJSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	insertUser(newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func login(c *gin.Context) {
	var usr user

	if err := c.BindJSON(&usr); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dbLogin(usr)
}
