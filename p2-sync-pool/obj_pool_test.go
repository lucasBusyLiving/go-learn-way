package main

import (
	"encoding/json"
	"sync"
	"testing"
)

// go test -bench . -benchmem
// Go 语言标准库内置的 testing 测试框架提供了基准测试(benchmark)的能力，能让我们很容易地对某一段代码进行性能测试。
type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
