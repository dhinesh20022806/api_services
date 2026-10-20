package middlewares

import (
	"net/http"
	"strings"

	"github.com/dhinesh20022806/api_s3/app/internal/grpcclient/iam"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {

	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokens := strings.Split(token, " ")

	user := iam.NewUser("", "", "", tokens[1])

	validUser, err := user.IsValidJwt()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "token invalid or expired",
		})

		return
	}
	ctx.Set("user", validUser)
	ctx.Next()
}
