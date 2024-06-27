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
	f := func(endNum int, gNum int) {
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
	startNum := 0
	go f(endNum, 1)
	go f(endNum, 2)

	ch <- startNum
	<-endCh
	close(ch)
	close(endCh)
}
