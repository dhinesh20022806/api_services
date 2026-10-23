// Package cmd Entry point of application
package cmd

import (
	"github.com/dhinesh20022806/api_s3/app/handlers"
	"github.com/dhinesh20022806/api_s3/app/middlewares"
	"github.com/gin-gonic/gin"
)

// RunServer is create http server
func RunServer() {

	server := gin.Default()

	server.POST("/login", handlers.Login)
	server.POST("/signup", handlers.SignUp)

	auth := server.Group("")

	server.MaxMultipartMemory = 4 << 30

	auth.Use(middlewares.Authenticate)

	{
		auth.POST("/bucket", handlers.POSTBucket)
		auth.PUT("/", handlers.PUTObject)
		auth.GET("/", handlers.GETObject)

	}

	err := server.Run(":8080")

	if err != nil {

		return
	}

}
