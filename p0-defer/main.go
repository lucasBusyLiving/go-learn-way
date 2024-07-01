package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("打印中", err)
		}
	}()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}
