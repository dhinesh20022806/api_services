// Package iam handling external service for Identity and Access Management
package iam

import (
	"context"

	genpb "github.com/dhinesh20022806/grpc_service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// User is Identity and Access Management structure
type User struct {
	Email    string
	password string
	Name     string
	jwt      string
	ID       string
}

var conn *grpc.ClientConn
var client genpb.AuthServiceClient

func init() {
	var err error

	conn, err = grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {

		panic(err)
	}

	client = genpb.NewAuthServiceClient(conn)
}

// NewUser is constructor of IamUser type
func NewUser(name, email, password, jwt string) *User {
	return &User{
		Email:    email,
		password: password,
		Name:     name,
		jwt:      jwt,
	}
}

// CreateUser used for singup it's from external services
func (i *User) CreateUser() (string, error) {
	ctx := context.Background()
	res, err := client.CreateUser(ctx, &genpb.CreateUserRequest{
		Email:    i.Email,
		Name:     i.Name,
		Password: i.password,
	})

	if err != nil {
		return "", err
	}
	return res.Jwt, nil
}

// Login is method execute on external services
func (i *User) Login() (string, error) {
	ctx := context.Background()

	res, err := client.IsValidUser(ctx, &genpb.IsValidUserRequest{
		Email:    i.Email,
		Password: i.password,
	})

	if err != nil {
		return "", err
	}

	return res.Jwt, nil

}

// IsValidJwt valid jwt token from external services
func (i *User) IsValidJwt() (User, error) {
	ctx := context.Background()
	user, err := client.IsValidJWT(ctx, &genpb.IsValidJWTRequest{
		Jwt: i.jwt,
	})

	if err != nil {

		return User{}, err
	}

	return User{
		Email: user.Email,
		ID:    user.Id,
		Name:  user.Name,
	}, nil
}
