package main

import (
	restApi "github.com/ljcnh/goAdPro/ch4/rest/restApi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

type RestServiceImpl struct{}

func (r *RestServiceImpl) Get(ctx context.Context, message *restApi.StringMessage) (*restApi.StringMessage, error) {
	return &restApi.StringMessage{Value: "Get hi:" + message.Value + "#"}, nil
}

func (r *RestServiceImpl) Post(ctx context.Context, message *restApi.StringMessage) (*restApi.StringMessage, error) {
	return &restApi.StringMessage{Value: "Post hi:" + message.Value + "@"}, nil
}
func main() {
	grpcServer := grpc.NewServer()
	restApi.RegisterRestServiceServer(grpcServer, new(RestServiceImpl))
	lis, _ := net.Listen("tcp", ":5000")
	grpcServer.Serve(lis)
}
