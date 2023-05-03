package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type application struct {
	ID      string `json:"id"`
	Company string `json:"companyId"`
	Title   string `json:"title"`
	Salary  int    `json:"salary"`
	Notes   string `json:"notes"`
}

var applications = []application{
	{ID: "0", Company: "Google", Title: "SD III", Salary: 140000, Notes: "fun place"},
}

func getApplications(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, applications)
}

func getApplicationById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range applications {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func createApplication(c *gin.Context) {
	var newAlbum application

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	applications = append(applications, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
