package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "用法：%s IP地址\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println(name, "不是一个有效的IP地址")
	} else {
		fmt.Println(name, "的IP地址为", addr)
	}
	os.Exit(0)
}

// douxiaobo@192 Golang % go run ip.go 123.123.0.1
// 123.123.0.1 的IP地址为 123.123.0.1
// douxiaobo@192 Golang % go run ip.go 百度
// 百度 不是一个有效的IP地址
// douxiaobo@192 Golang %
