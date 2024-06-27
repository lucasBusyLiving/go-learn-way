package main

import (
	"fmt"
)

// map panic 无法被 recover：fatal error: concurrent map writes。因为不是 panic，是 fatal
// recover只能捕获同一个协程中的panic，无法捕获其它协程的panic

func main() {
	// panic发生后 defer 会被执行完
	defer func() { fmt.Println("打印1") }()
	defer func() { fmt.Println("打印2") }()
	defer func() { fmt.Println("打印3") }()
	defer func() {
		// 不能对 recover 做函数封装
		// 源码实现角度：panic 发生后，就会对 defer 做检查，每个 defer是否有 recover，没有就继续抛出 panic。也就是 recover 的生效的作用域是
		// 当前函数
		// 设计角度：recover 作用于当前函数栈，子函数的 recover 管不了父函数。设计角度，一个函数的功能性应该是完整的。
		// recoverFromPanic应该也能不放在defer 中执行，在recoverFromPanic的上下文中，recover 没有在 defer 中执行，没有意义。
		recoverFromPanic()
	}()

	panic("触发异常 this panic can not be recovered")
}

func recoverFromPanic() {
	fmt.Println("recoverFromPanic can be execute")
	//
	if err := recover(); err != nil {
		fmt.Printf("Recovered from panic: %v\n", err)
	}
}
