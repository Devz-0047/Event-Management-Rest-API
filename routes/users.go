package routes

import (
	"net/http"

	"example.com/REST/models"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
)

func signup(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}

	if user.Email == "" || user.Password == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Email and password are required"})
		return
	}

	if err := user.Save(); err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
			context.JSON(http.StatusConflict, gin.H{"message": "Email already exists"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		}
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
func login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}

	err := user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}
