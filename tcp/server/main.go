package main

import (
	"fmt"
	"net" // net package provides a portable interface for network I/O, including TCP/IP, UDP, domain name system (DNS), and Unix domain sockets.
)

func main() {
	host := "127.0.0.1"
	port := "8080"                                  // 监听8080端口
	listen, err := net.Listen("tcp", host+":"+port) // 创建一个TCP监听器，用于监听TCP连接,开始监听指定的IP地址和端口号上的连接请求
	if err != nil {
		panic(err)
	}
	defer listen.Close() // 确保在程序退出前关闭监听器
	fmt.Println("Listening on", host+":"+port)
	for { // 循环不断地接受新的连接请求。
		conn, err := listen.Accept() //接受一个连接请求，如果有错误则处理错误，否则启动一个新的goroutine来处理这个连接。
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
			// panic(err)
		}
		// 处理连接，例如在一个单独的goroutine中
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// 在这里处理连接，读取和发送数据
	defer conn.Close() // 确保在函数退出前关闭连接
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer) // 从连接中读取数据		接收客户端发送的数据
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}
	data := buffer[:n]
	fmt.Println("Received data:", string(data)) // 打印接收到的数据		打印接收到的数据。

	// 发送HTTP响应
	response := "HTTP/1.1 200 OK\r\n\r\nHello World"        // 定义HTTP响应		构造 HTTP 响应字符串
	if _, err := conn.Write([]byte(response)); err != nil { //使用conn.Write()方法将HTTP响应发送给客户端。
		fmt.Println("Error writing response:", err)
		return
	}
	fmt.Println("Sent response:", response)
}
