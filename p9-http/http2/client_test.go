package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"golang.org/x/net/http2"
)

func TestHttp2(t *testing.T) {
	// 创建一个支持 HTTP/2 的传输层
	tr := &http2.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 在测试环境中忽略自签名证书的验证
		},
	}

	// 创建一个 HTTP 客户端并使用 HTTP/2 传输层
	client := &http.Client{Transport: tr}

	// 发起一个 HTTP/2 请求
	resp, err := client.Get("https://localhost:8082")
	if err != nil {
		log.Fatalf("Failed to get: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Printf("Response: %s\n", body)
}
