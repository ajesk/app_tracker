package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getApplications(c *gin.Context) {
	applications, err := queryAllApplications()

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	JSON(c, http.StatusCreated, applications)
}

func getApplicationById(c *gin.Context) {
	idStr := c.Param("applicationId")
	id, idErr := strconv.Atoi(idStr)

	if idErr != nil {
		fmt.Println(idErr.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	app, queryErr := queryApplicationById(id)

	if queryErr != nil {
		fmt.Println(queryErr.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.IndentedJSON(http.StatusOK, app)
}

func createApplication(c *gin.Context) {
	userIdStr := c.Param("userId")
	companyIdStr := c.Param("companyId")
	userId, userIdErr := strconv.Atoi(userIdStr)
	companyId, companyIdErr := strconv.Atoi(companyIdStr)

	if userIdErr != nil {
		fmt.Println(userIdErr.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if companyIdErr != nil {
		fmt.Println(companyIdErr.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var newApplication application
	if err := c.BindJSON(&newApplication); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := insertApplication(newApplication, companyId, userId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusCreated, id)
}
