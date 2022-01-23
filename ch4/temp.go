package main

import (
	"fmt"
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
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	doClientWork(client)
}
