// Package handlers end of the request cycle for handling Object
package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"uuid"

	"github.com/dhinesh20022806/api_s3/app/internal/grpcclient/dataroute"
	"github.com/dhinesh20022806/api_s3/app/internal/grpcclient/iam"
	"github.com/dhinesh20022806/api_s3/app/pkg/dbx"
	"github.com/gin-gonic/gin"
)

// RequestObject struct request body json data
type RequestObject struct {
	BucketID string `form:"bucketid"`
}

// PUTObject is store raw bytes to external services
func PUTObject(cxt *gin.Context) {

	var req RequestObject

	err := cxt.ShouldBind(&req)

	fmt.Printf("bucket id = %v", err)

	if err != nil {
		cxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "bad request",
		})
		return
	}

	file, err := cxt.FormFile("file")

	if err != nil {

		cxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
			"err":     err.Error(),
		})
		return
	}

	src, err := file.Open()

	if err != nil {
		cxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
			"err":     err.Error(),
		})
		return

	}
	defer func() {
		if err := src.Close(); err != nil {
			log.Printf("file close error: %v", err)
		}
	}()

	data, err := io.ReadAll(src)

	if err != nil {
		cxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
			"err":     err.Error(),
		})
		return
	}

	id := uuid.New()

	obj := dataroute.NewObject(id.String(), data)

	err = obj.PutObject()

	if err != nil {

		cxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
			"err":     err.Error(),
		})
		return
	}

	// meta data store

	err = dbx.InsertMetaData(id.String(), file.Filename, req.BucketID, "")

	if err != nil {

		cxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
		})
		return
	}

	cxt.JSON(http.StatusCreated, gin.H{
		"Message": "successfully created",
	})

}

// GETObject is serve bytes to actual filename how user store it
func GETObject(ctx *gin.Context) {

	var objectID string

	// id not email id

	val, exists := ctx.Get("user")

	if !exists {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return
	}

	user, ok := val.(*iam.User)

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return

	}

	err := ctx.ShouldBindBodyWithJSON(objectID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return
	}

	name, err := dbx.GetMyObject(objectID, user.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return
	}

	obj := dataroute.NewObject(objectID, []byte{})

	ObjectBytes, err := obj.GetObject()

	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return
	}

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	ctx.Header("Content-Type", "application/octet-stream")

	ctx.File(string(ObjectBytes))

}
