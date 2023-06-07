package main

import (
	"fmt"
	"time"
)

type Like struct {
	Count uint16
}

func (l *Like) Add(user string) {
	fmt.Printf("%s chang count: %d\n", user, l.Count)
	l.Count++
}

func (l *Like) Read(user string) {
	fmt.Printf("%s read count: %d\n", user, l.Count)
}

func ClickAdd(user string, like *Like) {
	for i := 0; i < 100; i++ {
		like.Add(user)
	}
}

func ClickRead(user string, like *Like) {
	for i := 0; i < 200; i++ {
		like.Read(user)
	}
}

func main() {
	like := new(Like)

	go ClickAdd("A", like)
	go ClickRead("B", like)
	go ClickAdd("C", like)
	go ClickRead("D", like)
	go ClickAdd("F", like)

	time.Sleep(1 * time.Second)
	fmt.Println(like.Count)
}
