package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 10; {
		i++
		fmt.Println(i)
	}
	a := 123
	fmt.Println(strconv.Itoa(a))
	ch := make(chan int)
	endCh := make(chan struct{})
	f := func(endNum int, gNum string) {
		for {
			select {
			case num, ok := <-ch:
				if !ok {
					fmt.Println("channel is closed")
					return
				}
				if num > endNum {
					endCh <- struct{}{}
					return
				}
				fmt.Println("channel", gNum, num)
				num += 1
				ch <- num
			}
		}
	}
	endNum := 100
	startNum := 1
	go f(endNum, "A")
	// 让 A 先打印
	ch <- startNum

	go f(endNum, "B")

	<-endCh
	close(ch)
	close(endCh)
}
