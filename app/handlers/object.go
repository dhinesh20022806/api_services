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

// PUTObject is store raw bytes to external services
func PUTObject(cxt *gin.Context) {

	var bucketID string

	err := cxt.ShouldBindBodyWithJSON(bucketID)

	if err != nil || len(bucketID) != 0 {
		cxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "bad request",
		})
		return
	}

	file, err := cxt.FormFile("file")

	if err != nil {

		cxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
		})
		return
	}

	src, err := file.Open()

	if err != nil {
		cxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
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
		})
		return
	}

	id := uuid.New()

	obj := dataroute.NewObject(id.String(), data)

	err = obj.PutObject()

	if err != nil {

		cxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
		})
		return
	}

	// meta data store

	err = dbx.InsertMetaData(id.String(), file.Filename, bucketID, "")

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
