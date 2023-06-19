package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetServerData(server string) string {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return fmt.Sprintf("from %s \n", server)
}

func Show(serverData ...interface{}) {
	fmt.Println(serverData...)
}

func main() {
	start := time.Now()

	getA := GetServerData("imA")
	getB := GetServerData("imB")
	getC := GetServerData("imC")
	Show(getA, getB, getC)
	fmt.Printf("get run : %s\n", time.Since(start))
}
