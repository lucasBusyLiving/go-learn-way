package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	// 设置HTTP处理程序
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, HTTP/3!")
	})

	// 加载自签名证书和私钥
	certFile := "server.crt"
	keyFile := "server.key"

	// 创建HTTP/3服务器
	server := &http3.Server{
		Addr:      ":4433",
		TLSConfig: generateTLSConfig(certFile, keyFile),
	}

	// 启动HTTP/3服务器
	log.Printf("Starting HTTP/3 server on https://localhost:4433")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start HTTP/3 server: %v", err)
	}
}

// 生成TLS配置
func generateTLSConfig(certFile, keyFile string) *tls.Config {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load key pair: %v", err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"h3", "http/1.1"},
	}
}
