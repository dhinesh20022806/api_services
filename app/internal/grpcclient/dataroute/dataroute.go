package dataroute

import (
	"context"

	genpb "github.com/dhinesh20022806/grpc_service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type Object struct {
	Id string
	Data []byte
}


var conn *grpc.ClientConn
var client genpb.ObjectServiceClient

func init() {
	var err error

	conn, err = grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	client = genpb.NewObjectServiceClient(conn)
}

func NewObject(id string, data []byte) *Object {
	return &Object{
		Id:id,
		Data: data,
	}
}

func (o *Object) PutObject() error {

	ctx := context.Background()

	res, err := client.StoreObject(ctx, &genpb.StoreObjectRequest{
		Data: o.Data,
	})

	if err != nil {
		return err
	}

	o.Id = res.ObjectId;
	
	return nil
}

func (o *Object) GetObject() ([]byte, error) {

	ctx := context.Background()

	res, err := client.GetObject(ctx, &genpb.GetObjectRequest{
		ObjectId: o.Id,
	});

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}