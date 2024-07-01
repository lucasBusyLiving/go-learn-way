package main

import (
	"fmt"
	"sync"
	"time"
)

/*
令牌桶实际实现：
1. 初始化 tokens 数量。桶的大小是 burst
2. 每次请求根据与上次请求的时间差（即惰性添加），去添加 token， 超过 burst 丢弃。然后减去本次需要的 token 数量
3. 每次请求需要加互斥锁
4. 核心思想就是如何添加令牌
*/

// TokenBucket 令牌桶结构体
type TokenBucket struct {
	sync.Mutex
	rate       int       // 每秒钟生成的令牌数量
	bucketSize int       // 令牌桶的容量
	tokens     int       // 当前桶中令牌的数量
	lastRefill time.Time // 上次补充令牌的时间
}

// NewTokenBucket 创建一个新的令牌桶
func NewTokenBucket(rate int, bucketSize int) *TokenBucket {
	return &TokenBucket{
		rate:       rate,
		bucketSize: bucketSize,
		tokens:     bucketSize,
		lastRefill: time.Now(),
	}
}

// refillTokens 补充令牌
func (tb *TokenBucket) refillTokens() {
	now := time.Now()
	duration := now.Sub(tb.lastRefill).Seconds()
	newTokens := int(duration * float64(tb.rate))
	if newTokens > 0 {
		tb.tokens = min(tb.bucketSize, tb.tokens+newTokens)
		tb.lastRefill = now
	}
}

// Allow 检查是否允许一个新的请求
func (tb *TokenBucket) Allow() bool {
	tb.Lock()
	defer tb.Unlock()

	tb.refillTokens()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func main() {
	rate := 5        // 每秒生成 5 个令牌
	bucketSize := 10 // 令牌桶的容量为 10
	limiter := NewTokenBucket(rate, bucketSize)

	var wg sync.WaitGroup

	allowChan := make(chan int, 20)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if limiter.Allow() {
				allowChan <- i
				fmt.Printf("Request %d: allowed\n", i)
			} else {
				fmt.Printf("Request %d: denied\n", i)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Allowed count", len(allowChan))

	time.Sleep(time.Second)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if limiter.Allow() {
				allowChan <- i
				fmt.Printf("Request %d: allowed\n", i)
			} else {
				fmt.Printf("Request %d: denied\n", i)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Allowed count", len(allowChan))

}
