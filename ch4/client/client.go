package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceClient struct {
	*rpc.Client
}

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

// 个人理解： 验证 保证HelloServiceClient 必有 Hello方法
// 删掉不影响输出
var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var replay string
	err = client.Hello("hello", &replay)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	fmt.Println(replay)
}

//func main() {
//	client, err := rpc.Dial("tcp", "localhost:1234")
//	if err != nil {
//		log.Fatal("dialing:", err)
//	}
//	var reply string
//	//err = client.Call("HelloService.Hello", "hello", &reply)
//	err = client.Call(HelloServiceName+".Hello", "hello", &reply)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(reply)
//}
