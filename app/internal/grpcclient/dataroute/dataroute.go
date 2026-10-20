// Package dataroute handling external services to store or get the object
package dataroute

import (
	"context"

	genpb "github.com/dhinesh20022806/grpc_service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Object which stored in external services
type Object struct {
	ID   string
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

// NewObject to access it's methods
func NewObject(ID string, data []byte) *Object {
	return &Object{
		ID:   ID,
		Data: data,
	}
}

// PutObject communicate with external service to store the object
func (o *Object) PutObject() error {

	ctx := context.Background()

	res, err := client.StoreObject(ctx, &genpb.StoreObjectRequest{
		Data: o.Data,
	})

	if err != nil {
		return err
	}

	o.ID = res.ObjectId

	return nil
}

// GetObject communicate with external services to get the object
func (o *Object) GetObject() ([]byte, error) {

	ctx := context.Background()

	res, err := client.GetObject(ctx, &genpb.GetObjectRequest{
		ObjectId: o.ID,
	})

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
