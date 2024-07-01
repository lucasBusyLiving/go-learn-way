package main

import "fmt"

//
// type Car struct {
// 	Name string
// }
//
// func main2() {
// 	incr := incrementer()
// 	x := incr()
// 	x++
// }
//
// func Speak() Car {
// 	bigCar := Car{"Big Car"}
// 	return bigCar
// }
//
// func incrementer() func() int {
// 	x := 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

func smallSlice() []int {
	slice := make([]int, 10) // 大数据量的 slice
	for i := range slice {
		slice[i] = i
	}
	return slice
}

func largeSlice() []int {
	slice := make([]int, 1000000) // 大数据量的 slice
	for i := range slice {
		slice[i] = i
	}
	return slice
}

func largeMap() map[int]int {
	m := make(map[int]int, 1000000) // 大数据量的 map
	for i := 0; i < 1000000; i++ {
		m[i] = i
	}
	return m
}

func main() {
	s := largeSlice()
	fmt.Println(s[999999])

	m := largeMap()
	fmt.Println(m[999999])

	ss := smallSlice()
	fmt.Println(ss)
}
