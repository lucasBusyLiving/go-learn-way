package main

import (
	"fmt"
)

type RingBuffer struct {
	data []byte
	head int
	tail int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		data: make([]byte, size),
	}
}

func (rb *RingBuffer) Write(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		rb.data[rb.head] = p[i]
		rb.head = (rb.head + 1) % len(rb.data)
		if rb.head == rb.tail {
			rb.tail = (rb.tail + 1) % len(rb.data) // 覆盖旧数据
		}
	}
	return len(p), nil
}

func (rb *RingBuffer) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p) && rb.tail != rb.head; i++ {
		p[i] = rb.data[rb.tail]
		rb.tail = (rb.tail + 1) % len(rb.data)
	}
	return len(p), nil
}

func main() {
	rb := NewRingBuffer(4)
	rb.Write([]byte{1, 2, 3, 4})
	buf := make([]byte, 4)
	rb.Read(buf)
	fmt.Println(buf) // 输出: [1 2 3 4]

	rb.Write([]byte{5, 6, 7, 8})
	rb.Read(buf)
	fmt.Println(buf) // 输出: [5 6 7 8]
}
