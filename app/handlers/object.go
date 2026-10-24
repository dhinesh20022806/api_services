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
func PUTObject(ctx *gin.Context) {

	bucketName := ctx.Param("bucketName")
	fileName := ctx.Param("fileName")

	file, err := ctx.FormFile("file")

	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
			"err":     err.Error(),
		})
		return
	}

	src, err := file.Open()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
			"err":     err.Error(),
		})
		return
	}

	id := uuid.New()

	obj := dataroute.NewObject(id.String(), data)

	err = obj.PutObject()

	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
			"err":     err.Error(),
		})
		return
	}

	// meta data store

	val, _ := ctx.Get("user")

	user := val.(iam.User)

	err = dbx.InsertMetaData(id.String(), fileName, bucketName, user.ID)

	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong!",
			"err":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "successfully created",
	})

}

// GETObject is serve bytes to actual filename how user store it
func GETObject(ctx *gin.Context) {

	bucketName := ctx.Param("bucketName")
	fileName := ctx.Param("fileName")

	val, exists := ctx.Get("user")

	if !exists {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return
	}

	user, ok := val.(iam.User)

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return

	}

	ObjectID, err := dbx.GetMyObject(fileName, bucketName, user.ID)

	if err != nil || ObjectID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found!", "msg": err.Error()})
		return
	}

	obj := dataroute.NewObject(ObjectID, []byte{})

	ObjectBytes, err := obj.GetObject()

	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found!", "err": err.Error()})
		return
	}

	ctx.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	ctx.Header("Content-Type", "application/octet-stream")

	ctx.Data(http.StatusOK, "application/octet-stream", ObjectBytes)

}
