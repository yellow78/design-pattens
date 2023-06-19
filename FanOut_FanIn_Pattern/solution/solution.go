package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 將多個輸入(channel)聚合
func Producer(servers ...string) <-chan string {
	producerCh := make(chan string, len(servers))

	go func() {
		defer close(producerCh)
		for _, serverName := range servers {
			producerCh <- serverName
		}
	}()

	return producerCh
}

// 執行併發任務
func Task(producerCh <-chan string) <-chan string {
	taskCh := make(chan string)

	go func() {
		defer close(taskCh)
		for serverName := range producerCh {
			taskCh <- GetServerData(serverName)
		}
	}()

	return taskCh
}

// 收集併發任務回傳之資訊
func Consumer(taskChs ...<-chan string) <-chan string {
	consumerCh := make(chan string, len(taskChs))

	var waitGroup sync.WaitGroup
	// 需等待任務之數量
	waitGroup.Add(len(taskChs))

	go func() {
		// 開始等待任務, 等到等待任務之數量為0時
		waitGroup.Wait()
		close(consumerCh)
	}()

	for _, task := range taskChs {
		go func(task <-chan string) {
			// 完成任務 wait -1數量
			defer waitGroup.Done()
			for response := range task {
				consumerCh <- response
			}
		}(task)
	}

	return consumerCh
}

func GetServerData(server string) string {
	getRand := rand.Intn(5)
	time.Sleep(time.Duration(getRand) * time.Second)
	return fmt.Sprintf("from %s : %d\n", server, getRand)
}

func Show(serverData ...interface{}) {
	fmt.Println(serverData...)
}

func main() {
	start := time.Now()
	producerCh := Producer("A", "B", "C", "D", "E")

	task1 := Task(producerCh)
	task2 := Task(producerCh)
	task3 := Task(producerCh)
	task4 := Task(producerCh)
	task5 := Task(producerCh)

	consumerCh := Consumer(task1, task2, task3, task4, task5)

	for new := range consumerCh {
		Show(new)
	}
	fmt.Printf("get run : %s", time.Since(start))
}
