package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("starting the thing")

	router := gin.Default()
	router.GET("/application", getApplications)
	router.POST("/application", createApplication)
	router.GET("/application/:id", getApplicationById)
	router.GET("/user", getUsers)
	router.GET("/user/:id", getUserById)
	router.POST("/user", createUser)

	router.Run("localhost:8080")
}
