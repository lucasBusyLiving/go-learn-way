package main

import (
	"context"
	"fmt"

	"golang.org/x/time/rate"
)

/*
令牌桶实际实现：
1. 初始化 tokens 数量。桶的大小是 burst
2. 每次请求根据与上次请求的时间差（即惰性添加），去添加 token， 超过 burst 丢弃。然后减去本次需要的 token 数量
3. 核心思想就是如何添加令牌
*/
func main() {
	r := rate.NewLimiter(10, 10)
	fmt.Println(r.Wait(context.Background()))
}
