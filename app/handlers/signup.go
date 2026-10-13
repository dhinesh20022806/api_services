package handlers

import (
	"net/http"

	grpcclient "github.com/dhinesh20022806/api_s3/app/internal/grpcclient/iam"
	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Name     string `binding:"required"`
}

func SignUp(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	newUser := grpcclient.NewUser(user.Name, user.Email, user.Password, "")

	token, err := newUser.CreateUser()

	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.SetCookie("token", token, 3600, "/", "localhost", false, false)

	ctx.JSON(http.StatusAccepted, gin.H{"message": "Your account created successfully!"})

}
