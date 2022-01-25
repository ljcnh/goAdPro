package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile(
		"server.crt", "server.grpc.io",
	)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
