package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MessageBox struct {
	message chan string
}

func (m *MessageBox) SendMessage(msg string) {
	m.message <- msg
}

func (m *MessageBox) ReadMessage() {
	for {
		fmt.Println(<-m.message)
	}
}

func CreatMessageBox() *MessageBox {
	msgBox := new(MessageBox)
	msgBox.message = make(chan string, 100)
	return msgBox
}

func main() {
	messageBox := CreatMessageBox()

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
