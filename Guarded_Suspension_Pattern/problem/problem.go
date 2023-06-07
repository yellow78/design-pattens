package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MessageBox struct {
	message []string
}

func (m *MessageBox) SendMessage(msg string) {
	m.message = append(m.message, msg)
}

func (m *MessageBox) ReadMessage() {
	for {
		fmt.Println(m.message[0])
		m.message = m.message[1:]
	}
}

func main() {
	messageBox := new(MessageBox)

	go func() {
		for {
			//隨機發送處理時間
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			messageBox.SendMessage("hot")
		}
	}()

	go func() {
		for {
			//隨機發送處理時間
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			messageBox.SendMessage("cool")
		}
	}()

	go func() {
		messageBox.ReadMessage()
	}()

	time.Sleep(10 * time.Second) //等待goroutine執行完畢
}
