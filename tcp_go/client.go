package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	host := "127.0.0.1"
	port := "8080"
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("请输入要发送的数据（输入'exit'退出）：")

	for {
		reader := bufio.NewReader(os.Stdin)
		sendData, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		sendData = sendData[:len(sendData)-1] // 移除末尾的换行符

		if sendData == "exit" {
			fmt.Println("退出程序。")
			break
		}

		encodedData := []byte(sendData)
		_, err = conn.Write(encodedData)
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}

		fmt.Println("消息发送成功。")
	}
}
