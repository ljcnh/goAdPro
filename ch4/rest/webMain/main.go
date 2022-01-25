package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	restApi "github.com/ljcnh/goAdPro/ch4/rest/restApi"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	err := restApi.RegisterRestServiceHandlerFromEndpoint(
		ctx, mux, "localhost:5000",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", mux)
}
