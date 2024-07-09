package main

import "fmt"

func main() {
	// type rune = int32
	// rune底层是 int32，也就是 unicode
	// '1' 单引号定义的类型为 rune
	s := "123"
	// s[0] 类型为 byte，这里 == 比较，go编译器自己做了类型转换
	fmt.Println(s[0] == '1') // s[0]->byte	'1' -> rune

	a := '1' // a 为 rune
	// s[0]->byte
	// fmt.Println(s[0] == a) // 自己显式定义的类型 go 不会做转换

	fmt.Println(s[0] == byte(a))

	b := '1' - '0' // b int32
	c := '1'
	d := '0'
	b = c - d
	println(b)
}
