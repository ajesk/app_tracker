package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getApplicationStatuses(c *gin.Context) {
	idStr := c.Param("applicationId")
	applicationId, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Printf("error processing application id: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	statuses, err := queryStatusesByApplication(applicationId)
	c.IndentedJSON(http.StatusOK, statuses)
}

func createStatus(c *gin.Context) {
	idStr := c.Param("applicationId")
	applicationId, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("error processing id for application %s: %s", idStr, err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var newStatus status
	if err := c.BindJSON(&newStatus); err != nil {
		fmt.Printf("error processing status object for application: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := insertApplicationStatus(newStatus, applicationId)

	if err != nil {
		fmt.Printf("error on insert: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// insertUser(newUser)
	c.IndentedJSON(http.StatusCreated, id)
}
