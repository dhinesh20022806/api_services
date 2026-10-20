// Package handlers response bucket curd operations
package handlers

import (
	"net/http"
	"strings"

	"github.com/dhinesh20022806/api_s3/app/internal/grpcclient/iam"
	"github.com/dhinesh20022806/api_s3/app/pkg/dbx"
	"github.com/gin-gonic/gin"
)

// POSTBucket used to create new Bucket
func POSTBucket(ctx *gin.Context) {

	var name string

	err := ctx.ShouldBindBodyWithJSON(name)

	if err != nil || len(strings.Trim(name, " ")) > 0 {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		return
	}

	val, exists := ctx.Get("user")

	if !exists {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})

		return
	}

	user, valid := val.(*iam.User)

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		return
	}

	bucket, err := dbx.CreateBucket(name, user.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "successfully Bucket Created!",
		"data": gin.H{
			"bucketId":   bucket.ID,
			"bucketName": bucket.BucketName,
		},
	})

}
