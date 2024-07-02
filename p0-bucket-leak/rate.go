package main

import (
	"fmt"
	"sync"
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

	rw := sync.RWMutex{}
	lock := sync.Mutex{}
	fmt.Println(rw, lock)
}

type myHeap []int

func (h myHeap) Len() int { return len(h) }

func (h myHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *myHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
