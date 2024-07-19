package main

import (
	"context"
)

func main() {
	// ctx: 1.参数传递，公参  2.协程间的同步
	ctx := context.Background()
	println(ctx)
}
