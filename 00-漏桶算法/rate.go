package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

/*
漏桶算法实现：
1. 漏桶思想是漏出的才会处理，也就是处理时间间隔是恒定的。后面进入桶内的请求需要等待
*/

func main() {
	// 每秒速率限制为 1000 个请求
	// 平均每个请求 1 ms
	rl := ratelimit.New(1000)

	prev := time.Now()
	for i := 0; i < 5; i++ {
		now := rl.Take()
		// 打印每个请求距离上个请求的时间
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
