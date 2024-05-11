package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 创建一个缓冲区来存储接收到的数据
	buffer := make([]byte, 1024)

	// 读取客户端发送的数据
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印请求（可选）
	fmt.Println("Received request:", string(buffer))

	// 响应客户端
	response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello, World!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// 创建一个TCP监听器，监听本地8080端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// 无限循环，等待客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 在一个新的goroutine中处理连接
		go handleConnection(conn)
	}
}

// Run the server: go run server2.go
// Received request: GET / HTTP/1.1
// Host: 127.0.0.1:8080
// Connection: keep-alive
// sec-ch-ua: "Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"
// sec-ch-ua-mobile: ?0
// sec-ch-ua-platform: "macOS"
// Upgrade-Insecure-Requests: 1
// User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36
// Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
// Sec-Fetch-Site: none
// Sec-Fetch-Mode: navigate
// Sec-Fetch-User: ?1
// Sec-Fetch-Dest: document
// Accept-Encoding: gzip, deflate, br, zstd
// Accept-Language: zh-CN,zh;q=0.9,es;q=0.8,en;q=0.7,ug;q=0.6

// Received request: GET /favicon.ico HTTP/1.1
// Host: 127.0.0.1:8080
// Connection: keep-alive
// sec-ch-ua: "Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"
// sec-ch-ua-mobile: ?0
// User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36
// sec-ch-ua-platform: "macOS"
// Accept: image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8
// Sec-Fetch-Site: same-origin
// Sec-Fetch-Mode: no-cors
// Sec-Fetch-Dest: image
// Referer: http://127.0.0.1:8080/
// Accept-Encoding: gzip, deflate, br, zstd
// Accept-Language: zh-CN,zh;q=0.9,es;q=0.8,en;q=0.7,ug;q=0.6

// From Terminal
// Received request: GET / HTTP/1.1
// Host: 127.0.0.1:8080
// User-Agent: curl/8.4.0
// Accept: */*

// From client.go
// Received request: GET / HTTP/1.1
// Host: 127.0.0.1:8080
