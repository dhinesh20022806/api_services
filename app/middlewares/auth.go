package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func Authenticate(ctx *gin.Context) {

	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// verify token grpc call

	
    ctx.Next()
}