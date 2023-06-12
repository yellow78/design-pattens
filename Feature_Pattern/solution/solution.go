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

func PushNew(msg string, getTime time.Time, p *PushBox) <-chan time.Duration {
	newsPushTime := make(chan time.Duration)

	go func() {
		time.Sleep(time.Duration(3 * time.Second))
		n := new(News)
		n.txt = msg
		n.getTime = getTime
		p.push = append(p.push, *n)
		runTime := time.Since(getTime)
		fmt.Printf("push run : %s\n", runTime)
		newsPushTime <- runTime
	}()

	return newsPushTime
}

func main() {
	pBox := new(PushBox)
	start := time.Now()

	allNews := []string{
		"ABC",
		"記得",
		"DEF～",
	}

	newsPushChannel := []<-chan time.Duration{}
	for _, news := range allNews {
		newsPushChannel = append(newsPushChannel, PushNew(news, start, pBox))
	}

	for index, newsCh := range newsPushChannel {
		fmt.Printf("news %d is push %s\n", index, <-newsCh)
	}

	time.Sleep(5)
	fmt.Println(pBox.push)
}
