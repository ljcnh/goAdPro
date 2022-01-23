package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const HelloServiceName3 = "HelloService"

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call(HelloServiceName3+".Login", "user:password", &reply)
	err = client.Call(HelloServiceName3+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
