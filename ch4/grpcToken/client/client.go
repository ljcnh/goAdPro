package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
func main() {
	port := "8080"
	auth := Authentication{
		User:     "gopher",
		Password: "password",
	}
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// ...
}
