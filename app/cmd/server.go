package cmd

import (
	"log"

	genpb "github.com/dhinesh20022806/grpc_service/proto/gen"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


var conn *grpc.ClientConn
var client *genpb.AuthServiceClient
var Client genpb.AuthServiceClient

func init() {
	addr := "localhost:50050";

	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		)

	if err != nil {
log.Fatalf("Failed to create connection: %v", err)
	}


	client := genpb.NewAuthServiceClient(conn)

	Client = client

}

func RunServer() {
	

	server := gin.Default()

	server.POST("/login", login)

	server.POST("/signup", signup)

	server.GET("/", getObject)
	server.PUT("/", putObject)

	server.Run(":8080")

	
   

	
	


}

func login(context *gin.Context) {
	
	

}

func signup(context *gin.Context) {


}

func getObject(context *gin.Context) {


}

func putObject(context *gin.Context) {



}