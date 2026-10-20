// Package handlers login request response
package handlers

import (
	"github.com/dhinesh20022806/api_s3/app/internal/grpcclient/iam"
	"github.com/gin-gonic/gin"
)

// LoginUser structure expected from request
type LoginUser struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Login controller create jwt token and stored to cookies
func Login(context *gin.Context) {

	var user LoginUser

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request data."})
		return
	}

	newUser := iam.NewUser("", user.Email, user.Password, "")

	token, err := newUser.Login()

	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid email or password."})
		return
	}

	context.SetCookie("token", token, 3600, "/", "localhost", false, false)

	context.JSON(200, gin.H{"message": "Login successful!"})

}
