// Package middlewares handles checks, append additional informations
package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dhinesh20022806/api_s3/app/internal/grpcclient/iam"
	"github.com/gin-gonic/gin"
)

// Authenticate is Middleware to check whether request is authenticate or not
// add additional information to rest of request flow
func Authenticate(ctx *gin.Context) {

	token := ctx.Request.Header.Get("Authorization")

	fmt.Printf("token=%s\n", token)
	if token == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokens := strings.Split(token, " ")

	user := iam.NewUser("", "", "", tokens[1])

	validUser, err := user.IsValidJwt()

	fmt.Printf("error = %v, validUser = %v", err, validUser)
	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "token invalid or expired",
		})

		return
	}
	ctx.Set("user", validUser)
	ctx.Next()
}
