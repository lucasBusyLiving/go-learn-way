package main

import (
	"fmt"
	"slices"
)

func foo() {
	type stu struct {
		Name string
		Age  int
	}
	m := map[stu]bool{}
	stu1 := stu{"stu1", 1}
	stu2 := stu{"stu1", 1}
	stu3 := stu{"stu1", 1}
	m[stu1] = true
	fmt.Println(m[stu2], m[stu3])
}

func main() {
	foo()
	slices.Max()
}
