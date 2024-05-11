package main

import (
	"fmt"
	"net"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 构建HTTP GET请求
	request := "GET / HTTP/1.1\r\nHost: 127.0.0.1:8080\r\n\r\n"

	// 发送请求到服务器
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取服务器的响应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印服务器的响应
	response := string(buffer[:n])
	fmt.Println(response)
}

// Terminal Output:
// curl http://127.0.0.1:8080
// Hello, World!%

// HTTP/1.1 200 OK
// Content-Type: text/plain

// Hello, World!
