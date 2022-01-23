package main

import (
	"log"
	"net"
	"net/rpc"
	_ "net/rpc/jsonrpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct {
	conn net.Conn
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// 支持tcp
func main() {
	RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		//go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		go rpc.ServeConn(conn)
	}
}

// 支持http
//func main() {
//	RegisterHelloService(new(HelloService))
//	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
//		var conn io.ReadWriteCloser = struct {
//			io.Writer
//			io.ReadCloser
//		}{
//			ReadCloser: r.Body,
//			Writer:     w,
//		}
//		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
//	})
//	http.ListenAndServe(":1234", nil)
//}

// main1	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
//		var conn io.ReadWriteCloser = struct {
//			io.Writer
//			io.ReadCloser
//		}{
//			ReadCloser: r.Body,
//			Writer:     w,
//		}
//		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
//	})
//	http.ListenAndServe(":1234", nil)

//func main() {
//	rpc.RegisterName("HelloService", new(HelloService))
//	listener, err := net.Listen("tcp", ":1234")
//	if err != nil {
//		log.Fatal("ListenTCP error:", err)
//	}
//	conn, err := listener.Accept()
//	if err != nil {
//		log.Fatal("Accept error:", err)
//	}
//	rpc.ServeConn(conn)
//}
