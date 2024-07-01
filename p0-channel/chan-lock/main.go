package main

import (
	"fmt"
	"time"
)

type chanLock struct {
	ch chan struct{}
}

func f(ch chan struct{}) {
	ch <- struct{}{}
	fmt.Println("进入")
	time.Sleep(time.Second * 2)
	defer func() { <-ch }()
}

func main() {
	var ch1 chan []int
	var ch2 chan map[int]int
	ch1 = make(chan []int, 10)
	ch2 = make(chan map[int]int, 10)
	ch1 <- make([]int, 10)
	ch2 <- map[int]int{1: 1}
	fmt.Println(<-ch1, <-ch2)

	cl := chanLock{make(chan struct{}, 1)}
	go f(cl.ch)
	go f(cl.ch)
	time.Sleep(time.Second * 5)

}
