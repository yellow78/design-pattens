package main

import (
	"fmt"
	"time"
)

type Like struct {
	Count uint16
}

func (l *Like) Add(user string) {
	l.Count++
}

// 每個使用者進入則點擊10000次
func ClickAdd(user string, like *Like) {
	// 測試10000, 數量太少不會有問題
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

	// 總計預期會50000但數量都不足
	time.Sleep(1 * time.Second)
	fmt.Println(like.Count)
}
