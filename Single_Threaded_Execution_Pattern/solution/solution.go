package main

import (
	"fmt"
	"sync"
	"time"
)

type Like struct {
	// 加入操控lock
	sync.Mutex
	Count uint16
}

func (l *Like) Add(user string) {
	// 線程鎖
	l.Lock()

	// 執行完解線程鎖
	defer l.Unlock()
	l.Count++
}

func ClickAdd(user string, like *Like) {
	for i := 0; i < 10000; i++ {
		like.Add(user)
	}
}

func main() {
	like := new(Like)

	go ClickAdd("A", like)
	go ClickAdd("B", like)
	go ClickAdd("C", like)
	go ClickAdd("D", like)
	go ClickAdd("F", like)

	time.Sleep(1 * time.Second)
	fmt.Println(like.Count)
}
