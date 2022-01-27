package main

import (
	"fmt"
	"sync"
)

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

func (l Lock) UnLock() {
	l.c <- struct{}{}
}

var counter int

// 注意 最终结果不一定是一样的
func main() {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				println("lock filed")
				return
			}
			counter++
			println("current counter", counter)
			l.UnLock()
		}()
	}
	wg.Wait()
	fmt.Print("counter  ")
	fmt.Println(counter)
}
