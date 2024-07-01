package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go func() {
		ch <- 2
		ch <- 2
		ch <- 2
		time.Sleep(2 * time.Second)
		ch <- 2
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	go func() {
		v, ok := <-ch
		fmt.Println(v, ok)
		v, ok = <-ch
		fmt.Println(v, ok)
		v, ok = <-ch
		fmt.Println(v, ok)
		v, ok = <-ch
		fmt.Println(v, ok)
		// 此时通道已经关闭
		v, ok = <-ch
		fmt.Println(v, ok)
		// 可以多次接收已经关闭的通道
		v, ok = <-ch
		fmt.Println(v, ok)
		// 不会进去循环
		for i := range ch {
			fmt.Println(i)
		}
	}()
	time.Sleep(5 * time.Second)
}
