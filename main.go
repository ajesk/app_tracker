package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("starting the thing")

	router := gin.Default()
	router.GET("/user/:userId/application", getApplications)
	router.POST("/user/:userId/company/:companyId/application", createApplication)
	router.GET("/application/:applicationId", getApplicationById)
	router.GET("/application/:applicationId/status", getApplicationStatuses)
	router.POST("/application/:applicationId/status", createStatus)
	router.GET("/user", getUsers)
	router.GET("/user/:userId", getUserById)
	router.POST("/user", createUser)
	router.GET("/company", getCompanies)
	router.GET("/company/:id", getCompanyById)
	router.POST("/company", createCompany)
	router.POST("/login", login)

	router.Run("localhost:8080")
}
