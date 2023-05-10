package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getCompanies(c *gin.Context) {
	companies, err := queryAllCompanies()
	if err != nil {
		fmt.Printf("error getting all companies: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, companies)
}

func getCompanyById(c *gin.Context) {
	idStr := c.Param("id")
	id, idErr := strconv.Atoi(idStr)

	if idErr != nil {
		fmt.Printf("error processing id %s", idErr.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	company, queryErr := queryCompanyById(id)

	if queryErr != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.IndentedJSON(http.StatusOK, company)
}

func createCompany(c *gin.Context) {
	var newCompany company

	if err := c.BindJSON(&newCompany); err != nil {
		return
	}

	id, err := insertCompany(newCompany)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusCreated, id)
}
