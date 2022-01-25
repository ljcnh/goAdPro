package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net/rpc"
	"time"
)

func doClientWork(client *rpc.Client) {
	helloCall := client.Go("path/to/pkg.HelloService.Hello", "hello", new(string), nil)

	time.Sleep(1 * time.Millisecond)
	// do some thing

	helloCall = <-helloCall.Done
	if err := helloCall.Error; err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	args := helloCall.Args.(string)
	reply := helloCall.Reply.(*string)
	fmt.Println(args, *reply)
}

func main() {
	s := grpc.NewServer()
	pb.RegisterYourOwnServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	s.Serve(lis)
	//port := "3999"
	//mux := http.NewServeMux()

	//creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//grpcServer := grpc.NewServer(grpc.Creds(creds))
	//
	//http.ListenAndServeTLS(port, "server.crt", "server.key",
	//	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		if r.ProtoMajor != 2 {
	//			mux.ServeHTTP(w, r)
	//			return
	//		}
	//		if strings.Contains(
	//			r.Header.Get("Content-Type"), "application/grpc",
	//		) {
	//			grpcServer.ServeHTTP(w, r) // gRPC Server
	//			return
	//		}
	//
	//		mux.ServeHTTP(w, r)
	//		return
	//	}),
	//)
	//mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//	fmt.Fprintln(w, "hello")
	//})
	//
	//http.ListenAndServeTLS(port, "server.crt", "server.key",
	//	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		mux.ServeHTTP(w, r)
	//		return
	//	}),
	//)

	//mux := http.NewServeMux()
	//h2Handler := h2c.NewHandler(mux, &http2.Server{})
	//server := &http.Server{Addr: ":3999", Handler: h2Handler}
	//server.ListenAndServe()
}
