package handlers

import (
	"io"
	"net/http"
	"uuid"

	"github.com/dhinesh20022806/api_s3/app/internal/grpcclient/dataroute"
	"github.com/dhinesh20022806/api_s3/app/internal/grpcclient/iam"
	"github.com/dhinesh20022806/api_s3/app/pkg/dbx"
	"github.com/gin-gonic/gin"
)

func PUTObject(cxt *gin.Context) {

	var bucketId string

	cxt.ShouldBindBodyWithJSON(bucketId)

	if len(bucketId) != 0 {
		cxt.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "bad request",
		})
		return
	}

	file, err := cxt.FormFile("file")

	if err != nil {

	}

	src, err := file.Open()

	if err != nil {

	}
	defer src.Close()

	data, err := io.ReadAll(src)

	id := uuid.New()

	obj := dataroute.NewObject(id.String(), data)

	err = obj.PutObject()

	if err != nil {

	}

	// meta data store

	err = dbx.InsertMetaData(id.String(), file.Filename, bucketId, "")

	if err != nil {

	}

	cxt.JSON(http.StatusCreated, gin.H{
		"Message": "successfully created",
	})

}

func GETObject(ctx *gin.Context) {

	var objectId string

	// id not email id

	val, exists := ctx.Get("user")

	if !exists {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return
	}

	user, ok := val.(*iam.IamUser)

	if !ok {

	}

}
