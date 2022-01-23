package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService1 struct{}

func (p *HelloService1) Hello(request *String, reply *String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func main() {
	rpc.RegisterName("HelloService1", new(HelloService1))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		//go rpc.ServeConn(conn)
	}
}
