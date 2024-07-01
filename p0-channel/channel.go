package main

import "fmt"

func recoverChanPanic() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ch := make(chan int)
	// 重复 close 会 panic
	close(ch)
	close(ch)
	return nil

}

func main2() {
	err := recoverChanPanic()
	fmt.Println(err)
}
