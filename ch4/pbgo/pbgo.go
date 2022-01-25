package main

import (
	"github.com/chai2010/pbgo/examples/hello.pb"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
)

func main() {
	router := hello_pb.HelloServiceHandler(new(HelloService))
	log.Fatal(http.ListenAndServe(":8080", router))
}

type HelloService struct{}

func (p *HelloService) Hello(request *hello_pb.String, reply *hello_pb.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}
func (p *HelloService) Echo(request *hello_pb.Message, reply *hello_pb.Message) error {
	*reply = *request
	return nil
}
func (p *HelloService) Static(request *hello_pb.String, reply *hello_pb.StaticFile) error {
	data, err := ioutil.ReadFile("./testdata/" + request.Value)
	if err != nil {
		return err
	}

	reply.ContentType = mime.TypeByExtension(request.Value)
	reply.ContentBody = data
	return nil
}
