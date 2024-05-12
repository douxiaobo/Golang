package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	host := "127.0.0.1"
	port := "8080"
	listen, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	fmt.Println("Listening on", host+":"+port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn) // 使用 goroutine 来处理每个连接
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭连接。")
			} else {
				fmt.Println("Error reading from connection:", err)
			}
			return
		}

		data := buffer[:n]
		fmt.Println("Received data:", string(data))

		// 发送响应
		response := "HTTP/1.1 200 OK\r\n\r\nHello World"
		if _, err := conn.Write([]byte(response)); err != nil {
			fmt.Println("Error writing response:", err)
			return
		}
		fmt.Println("Sent response:", response)
	}
}
