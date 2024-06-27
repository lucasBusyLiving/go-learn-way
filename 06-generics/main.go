package main

import (
	"fmt"
	"reflect"
)

// 泛型函数，适用于任意类型 T
func Map[T any](s []T, f func(T) T) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// 总结
// 泛型 提供编译时的类型安全，通过类型参数和类型约束确保在编译时捕获类型错误，避免运行时崩溃。
// any 提供灵活性，但缺乏编译时类型检查，需要运行时进行类型断言和检查，增加了运行时错误的风险和代码复杂性。
// 使用泛型时，类型安全在编译时得到保证，使得代码更加可靠和易于维护。
func main() {
	intSlice := []int{1, 2, 3}
	strSlice := []string{"a", "b", "c"}

	intResult := Map(intSlice, func(n int) int { return n * n })
	strResult := Map(strSlice, func(s string) string { return s + s })
	fmt.Println(reflect.TypeOf(intResult))
	fmt.Println(intResult) // Output: [1 4 9]
	fmt.Println(strResult) // Output: [aa bb cc]
}
