package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	host := "127.0.0.1"
	port := "8080"                              // 服务器监听的端口
	conn, err := net.Dial("tcp", host+":"+port) // 连接到服务器
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 使用conn进行读写操作
	// ...
	fmt.Println("请输入要发送的数据：")

	// 使用bufio.NewReader从标准输入读取一行数据
	// reader := bufio.NewReader(strings.NewReader(""))
	// 错误 "Error reading input: EOF" 源自于使用了 strings.Reader 来创建 bufio.Reader，实际上应该直接从标准输入读取数据。
	// strings.NewReader("") 创建了一个空的字符串读取器，当尝试从中读取数据时，立即到达EOF（文件结束符），因此产生了错误。
	reader := bufio.NewReader(os.Stdin)      // 提示用户输入要发送的数据
	sendData, err := reader.ReadString('\n') //读取一行直到遇到换行符，然后将末尾的换行符移除
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// 移除末尾的换行符
	sendData = sendData[:len(sendData)-1]

	// 发送数据，需要先将其编码为字节 slice
	encodedData := []byte(sendData)
	_, err = conn.Write(encodedData) //用户输入的数据编码为字节切片并发送到连接
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	// 接收数据
	fmt.Print("接收到的数据为：")
	buffer := make([]byte, 1024) //定义了一个缓冲区buffer来存储接收到的数据。
	n, err := conn.Read(buffer)  //使用conn.Read(buffer)从连接中读取数据，最多读取1024字节。
	if err != nil && err != io.EOF {
		fmt.Println("Error receiving data:", err)
		return
	}
	recvData := string(buffer[:n])
	fmt.Println(recvData) //将读取到的字节切片转换为字符串并打印出来
}
