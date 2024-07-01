package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	mu2 := sync.RWMutex{}
	fmt.Println(mu, mu2)
	one := sync.Once{}
	one.Do(func() {
	})

	// 获取当前的 GOMAXPROCS 值
	defaultGOMAXPROCS := runtime.GOMAXPROCS(0)
	fmt.Println("Default GOMAXPROCS:", defaultGOMAXPROCS)

	// 设置 GOMAXPROCS 值为 2
	runtime.GOMAXPROCS(2)

	// 获取新的 GOMAXPROCS 值
	newGOMAXPROCS := runtime.GOMAXPROCS(0)
	fmt.Println("New GOMAXPROCS:", newGOMAXPROCS)
}
