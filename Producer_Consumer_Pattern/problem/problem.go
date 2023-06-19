package main

import (
	"fmt"
	"time"
)

type User struct {
	id   uint32
	name string
}

// 訂單
var Jobs = []User{
	{1, "yellow"},
	{2, "red"},
	{3, "greed"},
}

// 新訂單進入
func Order(job chan<- User, id int, driverName string) {
	go Metch(Jobs[id], id, driverName)
}

// 客戶與司機配對完成
func Metch(user User, id int, driverName string) {
	fmt.Printf("consumer id: %d, %s get %s user\n", id, user.name, driverName)
}

func main() {
	send := make(chan User)

	// 只會為前三位司機接到單
	for i := 0; i < len(Jobs); i++ {
		name := fmt.Sprintf("imDriver%d", i)
		go Order(send, i, name)
	}

	time.Sleep(3 * time.Second)
}
