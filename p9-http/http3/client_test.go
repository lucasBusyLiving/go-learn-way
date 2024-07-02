package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/quic-go/quic-go/http3"
)

func TestHttp3(t *testing.T) {
	// 创建 HTTP/3 客户端
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 在测试环境中忽略自签名证书的验证
		},
	}
	client := &http.Client{
		Transport: roundTripper,
	}

	// 发起 HTTP/3 请求
	resp, err := client.Get("https://localhost:4433")
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
