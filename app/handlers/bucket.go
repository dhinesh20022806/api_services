// Package handlers response bucket curd operations
package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dhinesh20022806/api_s3/app/internal/grpcclient/iam"
	"github.com/dhinesh20022806/api_s3/app/pkg/dbx"
	"github.com/gin-gonic/gin"
)

// Request is Bucket request body struct
type Request struct {
	Name string `json:"name"`
}

// POSTBucket used to create new Bucket
func POSTBucket(ctx *gin.Context) {

	var req Request

	err := ctx.ShouldBindBodyWithJSON(&req)

	fmt.Printf("erro  = %v", err)

	if err != nil || len(strings.Trim(req.Name, " ")) == 0 {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
			"err":     err,
		})
		return
	}

	val, exists := ctx.Get("user")

	fmt.Printf("\nerro node error = %v, val = %v\n", exists, val)
	if !exists {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
			"err":     err,
		})

		return
	}

	user, valid := val.(iam.User)

	if !valid {
		fmt.Printf("\n erro node error = %v, user = %v \n", valid, user)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
			"err":     err,
		})
		return
	}

	bucket, err := dbx.CreateBucket(req.Name, user.ID)

	if err != nil {
		fmt.Printf("erro node error = %v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
			"err":     err,
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
