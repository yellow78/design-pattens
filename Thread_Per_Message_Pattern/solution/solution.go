package main

import (
	"fmt"
	"time"
)

type News struct {
	txt     string
	getTime time.Time
}

type PushBox struct {
	push []News
}

func PushNew(msg string, getTime time.Time, p *PushBox) {
	time.Sleep(time.Duration(3 * time.Second))
	n := new(News)
	n.txt = msg
	n.getTime = getTime
	p.push = append(p.push, *n)
	fmt.Printf("push run : %s\n", time.Since(getTime))
}

func main() {
	pBox := new(PushBox)
	start := time.Now()

	allNews := []string{
		"ABC",
		"記得",
		"DEF～",
	}

	for _, news := range allNews {
		go PushNew(news, start, pBox)
	}

	time.Sleep(10 * time.Second)
	fmt.Println(pBox.push)
}
