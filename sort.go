package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	arr := make([]int, 1001)
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		arr[i] = 1
	}
	cpuL := 8
	sliceLen := arrLen / cpuL
	t := reflect.TypeOf(sliceLen)
	fmt.Println(t)
	wg := sync.WaitGroup{}
	wg.Add(cpuL)
	ansChan := make(chan int, cpuL)
	for i := 0; i < cpuL; i++ {
		start := i * sliceLen
		end := start + sliceLen
		if i == cpuL-1 {
			end = arrLen
		}
		go func() {
			defer wg.Done()
			sumRes := 0
			for m := start; m < end; m++ {
				sumRes += arr[m]
			}
			ansChan <- sumRes
		}()
	}
	wg.Wait()

	close(ansChan)

	ans := 0
	for i := range ansChan {
		ans += i
	}
	for i := range ansChan {
		fmt.Println(i, "close")
	}
	fmt.Println("done", ans)
}
