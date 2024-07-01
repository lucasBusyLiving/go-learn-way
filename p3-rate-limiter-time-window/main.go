package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter 是一个固定窗口限流器
type RateLimiter struct {
	sync.Mutex
	interval    time.Duration // 时间窗口
	maxRequests int           // 时间窗口内允许的最大请求数
	requests    int           // 当前窗口内的请求数
	windowStart time.Time     // 当前时间窗口的起始时间
}

// NewRateLimiter 创建一个新的 RateLimiter
func NewRateLimiter(interval time.Duration, maxRequests int) *RateLimiter {
	return &RateLimiter{
		interval:    interval,
		maxRequests: maxRequests,
		windowStart: time.Now(),
	}
}

// Allow 检查是否允许一个新的请求
func (rl *RateLimiter) Allow() bool {
	rl.Lock()
	defer rl.Unlock()

	now := time.Now()

	// 如果当前时间超过了时间窗口，则重置窗口
	if now.Sub(rl.windowStart) > rl.interval {
		rl.windowStart = now
		rl.requests = 0
	}

	// 检查当前窗口内的请求数是否超过了限制
	if rl.requests < rl.maxRequests {
		rl.requests++
		return true
	}

	return false
}

func main() {
	limiter := NewRateLimiter(1*time.Second, 5) // 每秒最多允许5个请求

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if limiter.Allow() {
				fmt.Printf("Request %d: allowed\n", i)
			} else {
				fmt.Printf("Request %d: denied\n", i)
			}
		}(i)
	}

	wg.Wait()
}
