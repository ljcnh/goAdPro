package main

import (
	"sync"
)

// 全局变量
var counter int

func main() {
	var wg sync.WaitGroup
	var l sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			l.Lock()
			defer l.Unlock()
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	println(counter)
}
