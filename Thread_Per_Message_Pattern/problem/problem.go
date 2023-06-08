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

func (p *PushBox) PushNew(msg string, getTime time.Time) {
	time.Sleep(time.Duration(3 * time.Second))
	n := new(News)
	n.txt = msg
	n.getTime = getTime
	p.push = append(p.push, *n)
	fmt.Println(p.push)
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
		pBox.PushNew(news, start)
	}

	fmt.Printf("push run : %s\n", time.Since(start))
}
