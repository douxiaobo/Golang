package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " <port>")
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp4", nil, udpAddr)
	checkError(err)
	_, err = conn.Write([]byte("Hello, world!"))
	checkError(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(string(buf[0:n]))
	conn.Close()
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

// douxiaobo@192 UDP Socket % go run client.go localhost:1200
// 2024-05-31 19:44:30.126111 +0800 CST m=+28.374626918
// douxiaobo@192 UDP Socket % go run client.go 0.0.0.0:1200
// 2024-05-31 19:45:10.941369 +0800 CST m=+69.189653001
// douxiaobo@192 UDP Socket %
// douxiaobo@192 UDP Socket % go run client.go localhost:1200
// 2024-05-31 19:45:40.471063 +0800 CST m=+3.768684168
// douxiaobo@192 UDP Socket % go run client.go 127.0.0.1:1200
// 2024-05-31 19:45:51.792363 +0800 CST m=+15.089920376
