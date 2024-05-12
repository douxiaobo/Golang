package main

import (
	"fmt"
	"net"
	"sync"
)

func getString(socket *net.UDPConn) string {
	buffer := make([]byte, 65536)
	size, _, err := socket.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("读取数据时发生错误:", err)
		return ""
	}
	return string(buffer[:size])
}

func main() {
	var wg sync.WaitGroup // 声明一个 sync.WaitGroup 对象，用来等待两个并发任务（goroutines）完成。WaitGroup 是一个同步原语，用于等待一组goroutines完成它们的工作。
	wg.Add(2)             // 等待两个goroutine完成 用来初始化一个 sync.WaitGroup 对象，表示有两个并发任务（goroutines）将要执行。

	// WaitGroup 有两件事情需要完成：一个是在 socketA 上发送和接收数据，另一个是在 socketB 上发送和接收数据。
	// 因此，我们需要创建两个 goroutine，每个 goroutine 负责一个套接字。

	// 以下代码创建两个UDP套接字，并分别向对方发送消息，接收对方的消息并打印。

	// 创建套接字 A
	socketA, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 50000})
	if err != nil {
		fmt.Println("监听套接字 A 时发生错误:", err)
		return
	}
	defer socketA.Close()

	// 创建套接字 B
	socketB, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 50001})
	if err != nil {
		fmt.Println("监听套接字 B 时发生错误:", err)
		return
	}
	defer socketB.Close()

	go func() {
		defer wg.Done()
		_, err := socketA.WriteToUDP([]byte("This is A"), &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 50001})
		if err != nil {
			fmt.Println("套接字 A 发送数据时发生错误:", err)
			return
		}
		message := getString(socketA)
		fmt.Println("来自 B 的消息:", message)
	}()

	go func() {
		defer wg.Done()
		_, err := socketB.WriteToUDP([]byte("This is B"), &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 50000})
		if err != nil {
			fmt.Println("套接字 B 发送数据时发生错误:", err)
			return
		}
		message := getString(socketB)
		fmt.Println("来自 A 的消息:", message)
	}()

	wg.Wait()
}

// 这段Go代码实现了与Rust代码相同的功能：两个UDP套接字相互发送消息，并打印接收到的消息。
// 使用了sync.WaitGroup来等待两个goroutine完成。
// 注意，Go的net.UDPConn提供了ReadFromUDP和WriteToUDP方法来处理UDP数据的读写。
// 此外，Go的错误处理通常通过检查返回的错误值来进行。
