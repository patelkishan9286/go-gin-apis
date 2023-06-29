package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	ID        string `json:"id"`
}

var users = []User{
	{
		Email:     "john@example.com",
		FirstName: "John",
		ID:        "1",
	},
	{
		Email:     "jane@example.com",
		FirstName: "Jane",
		ID:        "2",
	},
	{
		Email:     "alice@example.com",
		FirstName: "Alice",
		ID:        "3",
	},
	{
		Email:     "bob@example.com",
		FirstName: "Bob",
		ID:        "4",
	},
	{
		Email:     "emma@example.com",
		FirstName: "Emma",
		ID:        "5",
	},
}

func main() {

	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/user/:id", getUserById)
	router.POST("/add", addUser)
	router.PUT("/update/:id", updateUser)

	router.Run(":3001")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Users retrieved",
		"success": true,
		"users":   users,
	})
}

// POST /add endpoint
func addUser(c *gin.Context) {

	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON provided", "error": err.Error(), "success": false})
		return
	}

	// Check if user already exists
	for _, user := range users {
		if user.Email == newUser.Email {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "User already exists", "success": false})
			return
		}
	}

	newUser.ID = strconv.Itoa(len(users) + 1)
	users = append(users, newUser)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "User added",
		"success": true,
	})
}

// GET /user/:id endpoint
func getUserById(c *gin.Context) {

	id := c.Param("id")

	for _, user := range users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{
				"success": true,
				"user":    user,
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "User not found",
		"success": false,
	})
}

// PUT /update/:id endpoint

func updateUser(c *gin.Context) {

	id := c.Param("id")
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON provided",
			"error":   err.Error(),
			"success": false,
		})
	}

	for i, user := range users {
		if user.ID == id {
			users[i].Email = newUser.Email
			users[i].FirstName = newUser.FirstName

			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "User updated",
				"success": true,
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found", "success": false})

}
