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
func Order(getJob chan<- User, id int) {
	getJob <- Jobs[id]
}

// 客戶與司機配對完成
func Metch(getJob <-chan User, id int, driverName string) {
	for userInfo := range getJob {
		fmt.Printf("consumer id: %d, %s get %s user\n", id, userInfo.name, driverName)
	}
}

func main() {
	sendjob := make(chan User)

	// 五位司機線上處理訂單
	driverCount := 5

	for i := 0; i < len(Jobs); i++ {
		go Order(sendjob, i)
	}

	for i := 0; i < driverCount; i++ {
		name := fmt.Sprintf("imDriver%d", i)
		go Metch(sendjob, i, name)
	}

	time.Sleep(2 * time.Second)
}
