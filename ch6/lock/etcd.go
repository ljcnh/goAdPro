package main

import (
	"github.com/zieckey/etcdsync"
	"log"
)

// 先检查/lock路径下是否有值，如果有值，说明锁已经被别人抢了
// 如果没有值，那么写入自己的值。写入成功返回，说明加锁成功。写入时如果节点被其它节点写入过了，那么会导致加锁失败，这时候到 3
// watch /lock下的事件，此时陷入阻塞
// 当/lock路径下发生事件时，当前进程被唤醒。检查发生的事件是否是删除事件（说明锁被持有者主动unlock），或者过期事件（说明锁过期失效）。如果是的话，那么回到 1，走抢锁流程。
func main() {
	m, err := etcdsync.New("/lock", 10, []string{"http://127.0.0.1:2379"})
	if m == nil || err != nil {
		log.Printf("etcdsync.New failed")
		return
	}
	err = m.Lock()
	if err != nil {
		log.Printf("etcdsync.Lock failed")
		return
	}
	log.Printf("etcdsync.Lock OK")
	log.Printf("Get the lock. Do something here.")
	err = m.Unlock()
	if err != nil {
		log.Printf("etcdsync.Unlock failed")
	} else {
		log.Printf("etcdsync.Unlock OK")
	}
}
