package main

import "fmt"

// map panic 无法被 recover
// recover只能捕获同一个协程中的panic，无法捕获其它协程的panic

func main() {
	defer func() {
		// 不能对 recover 做函数封装
		// 源码实现角度：panic 发生后，就会对 defer 做检查，是否有 recover，没有就继续抛出 panic。不可能再去检查子函数，如果
		// 子函数的 recover 潜逃了很多层，还需要一个个去检查，成本太高
		// 设计角度：recover 作用于当前函数栈，子函数的 recover 管不了父函数。设计角度，一个函数的功能性应该是完整的。
		// recoverFromPanic应该也能不放在defer 中执行，在recoverFromPanic的上下文中，recover 没有在 defer 中执行，没有意义。
		recoverFromPanic()
	}()

	panic("this panic can not be recovered")
}

func recoverFromPanic() {
	if err := recover(); err != nil {
		fmt.Printf("Recovered from panic: %v\n", err)
	}
}
