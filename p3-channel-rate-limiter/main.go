package main

import (
	"fmt"
	"time"
)

// 限流器结构体
type Limiter struct {
	rate   int           // 每秒允许的请求数量
	burst  int           // 允许的突发请求数量
	tokens chan struct{} // 令牌通道
}

// 创建限流器
func NewLimiter(rate int, burst int) *Limiter {
	l := &Limiter{
		rate:   rate,
		burst:  burst,
		tokens: make(chan struct{}, burst),
	}

	// 初始化令牌通道
	for i := 0; i < burst; i++ {
		l.tokens <- struct{}{}
	}

	// 定时补充令牌
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rate))
		for range ticker.C {
			select {
			case l.tokens <- struct{}{}:
				fmt.Println("fill token done")
			default:
				fmt.Println("too many tokens")
			}
		}
	}()

	return l
}

// 检查是否允许请求
func (l *Limiter) Allow() bool {
	select {
	case <-l.tokens:
		return true
	default:
		return false
	}
}

func main() {
	limiter := NewLimiter(2, 5) // 每秒 2 个请求，突发 5 个

	for i := 1; i <= 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d is allowed\n", i)
		} else {
			fmt.Printf("Request %d is rejected\n", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
	time.Sleep(time.Second * 5)
}
