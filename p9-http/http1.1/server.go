package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond) // 模拟一些处理时间
	fmt.Fprintf(w, "Hello, HTTP/1.1!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting HTTP/1.1 server on :8081")
	http.ListenAndServe(":8081", nil)
}
