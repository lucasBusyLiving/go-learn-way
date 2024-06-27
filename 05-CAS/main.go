package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var value int32 = 0

	// 启动一个 goroutine 不断尝试修改值
	go func() {
		for i := 10; i < 100; i++ {
			//
			atomic.CompareAndSwapInt32(&value, 0, int32(i))
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// 主 goroutine 每隔一段时间打印值
	for {
		fmt.Println("Value:", value)
		time.Sleep(time.Second)
	}
}
