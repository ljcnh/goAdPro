package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

//type KVStoreServiceInterface = interface {
//	Get(key string, value *string) error
//	Set(kv [2]string, reply *struct{}) error
//	Watch(timeoutSecond int, keyChanged *string) error
//}

func main() {
	rpc.RegisterName("KVStoreService", NewKVStoreService())
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		fmt.Println(conn)
		rpc.ServeConn(conn)
	}
}

type KVStoreService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}
	return fmt.Errorf("not found")
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	key, value := kv[0], kv[1]
	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}
	p.m[key] = value
	return nil
}

func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	fmt.Println(id)

	ch := make(chan string, 10)
	p.mu.Lock()
	p.filter[id] = func(key string) {
		ch <- key
	}
	p.mu.Unlock()
	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		fmt.Println("Watch 1")
		return fmt.Errorf("timeout")
	case key := <-ch:
		fmt.Println("Watch 2")
		*keyChanged = key
		return nil
	}
	return nil
}
