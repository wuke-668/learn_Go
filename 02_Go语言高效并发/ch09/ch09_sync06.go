package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum   int
	mutex sync.RWMutex
)

func main()  {
	run()
}

func run() {
	for i := 0; i < 100; i++ {
		go add(10)
	}
	for i:=0;i<10;i++{
		go fmt.Println("和为 ",readSum())
	}
	// 防止提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("和为: ", sum)
}

func add(i int) {
	mutex.Lock()
	sum += i
	mutex.Unlock()
}
func readSum() int {
	mutex.RLock()
	defer mutex.RUnlock()
	b:=sum
	return b
}