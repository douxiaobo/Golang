package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 创建一个UDP连接
	conn, err := net.Dial("udp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 读取用户输入的摄氏温度
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter temperature in Celsius: ")
	temperatureStr, _ := reader.ReadString('\n')
	temperatureStr = strings.TrimSpace(temperatureStr) // 去除换行符

	// 将摄氏温度转换为浮点数
	temperature, err := strconv.ParseFloat(temperatureStr, 64)
	if err != nil {
		fmt.Println("Error parsing temperature:", err)
		os.Exit(1)
	}

	// 如果需要，可以将temperature格式化为两位小数后再发送（虽然这通常不是必要的，因为服务器会进行转换）
	formattedTemperature := fmt.Sprintf("%.1f", temperature)

	// 将温度编码为字节并发送给服务器
	// _, err = conn.Write([]byte(strconv.FormatFloat(temperature, 'f', -1, 64)))
	// 将格式化的温度编码为字节并发送给服务器
	_, err = conn.Write([]byte(formattedTemperature))
	if err != nil {
		fmt.Println("Error sending temperature:", err)
		os.Exit(1)
	}

	// 读取服务器的响应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving:", err)
		os.Exit(1)
	}

	// 打印服务器的响应
	fmt.Println(string(buffer[:n]))
}
