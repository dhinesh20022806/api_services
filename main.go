package main

import "github.com/gin-gonic/gin"



func main() {


	server := gin.Default()

	server.GET("/", getObject)
	server.PUT("/", putObject)

	server.Run(":8080")
}

func getObject(context *gin.Context) {


}

func putObject(context *gin.Context) {

}