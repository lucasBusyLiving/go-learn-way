package main

import (
	"fmt"
	"strings"
)

func str() {
	s1 := "abc"
	s2 := "b"
	res := strings.Contains(s1, s2)
	fmt.Println(res)
}
