package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 创建一个UDP连接
	conn, err := net.ListenPacket("udp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Binding UDP to port 8888")

	buffer := make([]byte, 1024)

	// 无限循环以接收数据
	for {
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		// 尝试将 addr 断言为 *net.UDPAddr
		udpAddr, ok := addr.(*net.UDPAddr)
		if !ok {
			// 如果 addr 不是 *net.UDPAddr 类型，处理错误或采取其他措施
			fmt.Println("Error casting addr to *net.UDPAddr")
			continue
		}

		// 现在你可以安全地访问 udpAddr.Port
		fmt.Printf("Received from %s:%d\n", udpAddr.IP, udpAddr.Port)

		// 转换接收到的数据为字符串并尝试将其解析为浮点数
		dataStr := string(buffer[:n])
		celsius, err := strconv.ParseFloat(strings.TrimSpace(dataStr), 64)
		if err != nil {
			fmt.Printf("Error parsing temperature: %v\n", err)
			continue
		}

		// 将摄氏度转换为华氏温度
		fahrenheit := celsius*1.8 + 32

		// 构造回复消息
		sendData := fmt.Sprintf("转换后的温度（单位：华氏温度）：%.1f\n", fahrenheit)

		// 发送回复
		_, err = conn.WriteTo([]byte(sendData), addr)
		if err != nil {
			fmt.Println("Error writing:", err)
			continue
		}

		fmt.Printf("Received from %s:%d. Sent: %s\n", addr.Network(), udpAddr.Port, sendData)
		break
	}
}
