package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User Blueprint
type User struct {
	Name  string `json:"name" binding:"required,min=2,max=50"`
	Email string `json:"email" binding:"required,email"`
}

// In-Memory Slice to Store Users
var users = []User{}

func main() {
	// Router, that Handles Requests
	r := gin.Default()

	// Handler that Receives Data from Requests & Store in "users" DB
	r.POST("/users", func(ctx *gin.Context) {
		var user User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Validation Failed",
			})
			return
		}

		users = append(users, user)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "User Record Created",
		})
	})

	// Handler that shows List of Users in our DB
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	// Command to Run the Router/Server
	r.Run(":8080")
}
