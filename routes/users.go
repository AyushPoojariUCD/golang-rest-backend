package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-rest-backend/utils"
	"go-rest-backend/models"
)

// Signup User
func signUp(c *gin.Context) {

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

// Login User
func login(c *gin.Context) {

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidateCredential()
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"token": token,
	})
}

// Get all users
func getUsers(c *gin.Context) {

	users, err := models.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not retrieve users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Get user by ID
func getUserByID(c *gin.Context) {

	id := c.Param("id")

	user, err := models.GetUserByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update User
func updateUser(c *gin.Context) {

	id := c.Param("id")

	user, err := models.GetUserByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	var updatedUser models.User

	err = c.ShouldBindJSON(&updatedUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedUser.ID = user.ID

	err = updatedUser.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}

// Delete User
func deleteUser(c *gin.Context) {

	id := c.Param("id")

	err := models.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not delete user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}