package handlers

import (
	"github.com/gin-gonic/gin"
)

// GetUsers handles the GET request to retrieve users
func GetUsers(c *gin.Context) {
	// Logic to retrieve users from the database or any other source
	users := []string{"Alice", "Bob", "Charlie"}

	// Respond with the list of users
	c.JSON(200, gin.H{
		"users": users,
	})
}

func NewUser(c *gin.Context) {

	var newUserRequest NewUserRequest

	// Bind the JSON request body to the NewUserRequest struct
	if err := c.ShouldBindJSON(&newUserRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Respond with the newly created user
	c.JSON(201, gin.H{
		"user": newUserRequest,
	})
}
