package cmd

import (
	"github.com/dhinesh20022806/api_s3/app/handlers"
	"github.com/dhinesh20022806/api_s3/app/middlewares"
	"github.com/gin-gonic/gin"
)



func RunServer() {

	server := gin.Default()

	server.POST("/login", handlers.Login)
	server.POST("/signup", handlers.SignUp)

	auth := server.Group("")

	server.MaxMultipartMemory = 4 << 30

	auth.Use(middlewares.Authenticate)

	{
		auth.PUT("/", handlers.PUTObject)
		auth.GET("/", )

	}


	server.Run(":8080")

}
