package main

import (
	"fmt"
	"time"
)

func leakyFunction(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Goroutine exiting")
			return
		default:
			// 模拟一些工作
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	done := make(chan bool)

	// 启动一个goroutine，但没有任何机制来通知它退出
	go leakyFunction(done)

	// 主程序等待一段时间后退出
	time.Sleep(3 * time.Second)
	fmt.Println("Main function exiting")
}
