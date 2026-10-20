package handlers

import (
	"net/http"
	"strings"

	"github.com/dhinesh20022806/api_s3/app/pkg/dbx"
	"github.com/gin-gonic/gin"
)

type Bucket struct {
	Id   string
	name string
}

func POSTBucket(ctx *gin.Context) {

	var name string

	ctx.ShouldBindBodyWithJSON(name)

	if len(strings.Trim(name, " ")) > 0 {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		return

	}

	dbx.CreateBucket(name)

}
