package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 创建一个简单的处理器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, HTTP/2!")
	})

	// 加载自签名证书和私钥
	certFile := "server.crt"
	keyFile := "server.key"

	// 创建一个 HTTPS 服务器
	server := &http.Server{
		Addr: ":8082",
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{loadTLSConfig(certFile, keyFile)},
			NextProtos:   []string{"h2", "http/1.1"},
		},
	}

	// 启动 HTTPS 服务器
	log.Printf("Starting HTTP/2 server on https://localhost:8080")
	if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
		log.Fatalf("Failed to start HTTP/2 server: %v", err)
	}
}

// 加载 TLS 配置
func loadTLSConfig(certFile, keyFile string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load key pair: %v", err)
	}
	return cert
}
