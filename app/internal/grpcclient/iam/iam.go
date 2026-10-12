package grpcclient

import (
	"context"

	genpb "github.com/dhinesh20022806/grpc_service/proto/gen"
	"google.golang.org/grpc"
)

type IamUser struct {
	email    string
	password string
	name     string
	jwt      string
	valid    bool
}

var conn *grpc.ClientConn
var client genpb.AuthServiceClient

func init() {
	var err error

	conn, err = grpc.NewClient("localhost:50050")

	if err != nil {

	}

	client = genpb.NewAuthServiceClient(conn)
}

func NewUser(name, email, password, jwt string) *IamUser {
	return &IamUser{
		email:    email,
		password: password,
		name:     name,
		jwt:      jwt,
	}
}

func (i *IamUser) CreateUser() (string, error) {
	ctx := context.Background()
	res, err := client.CreateUser(ctx, &genpb.CreateUserRequest{
		Email:    i.email,
		Name:     i.name,
		Password: i.password,
	})

	if err != nil {
		return "", err
	}
	return res.Jwt, nil
}

func (i *IamUser) Login() (string, error) {
	ctx := context.Background()

	res, err := client.IsValidUser(ctx, &genpb.IsValidUserRequest{
		Email:    i.email,
		Password: i.password,
	})

	if err != nil {
		return "", err
	}

	return res.Jwt, nil

}

func (i *IamUser) IsValidJwt() (bool, error) {
	ctx := context.Background()
	_, err := client.IsValidJWT(ctx, &genpb.IsValidJWTRequest{
		Jwt: i.jwt,
	})

	if err != nil {

		return false, err
	}
	return true, nil
}
