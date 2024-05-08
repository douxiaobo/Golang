package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}
	// fmt.Println("Server running on", host)

	port := 12345

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listen.Close()

	// lnAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", host, port))
	// if err != nil {
	// 	fmt.Println("Error resolving address:", err)
	// 	return
	// }

	// listener, err := net.ListenTCP("tcp", lnAddr)
	// if err != nil {
	// 	fmt.Println("Error listening:", err)
	// 	return
	// }
	// defer listener.Close()

	// listener.SetDeadline(time.Now().Add(time.Second)) // 设置超时时间以限制 Backlog 大小为 1
	// defer listener.SetDeadline(time.Time{})           // 移除超时限制

	fmt.Println("Waiting for connections...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
			// return
		}
		go handleConnection(conn)
		// break	//进不入Connected by...
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Connected by %s\n", conn.RemoteAddr().String())

	// 在这里添加处理连接的具体逻辑
	for {
		//接收
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			conn.Close()
			return
		}
		msg := string(buffer[:n])
		// fmt.Printf("Received: %s", msg)
		// reader := bufio.NewReader(conn)
		// msg, err := reader.ReadString('\n')
		// if err != nil {
		// 	fmt.Println("Error reading:", err)
		// 	conn.Close()
		// 	return
		// }
		// msg = strings.TrimSpace(msg)
		fmt.Printf("Message Received from client: %s\n", msg)

		if msg == "byebye\n" { //接收了，其实上退出了，但是在Terminal上还没退出
			fmt.Println("Client said goodbye.")
			break
		}

		//发送
		fmt.Print("Enter message to send to cient (type 'byebye' to exit):")
		send_data := bufio.NewReader(os.Stdin)
		sendData, err := send_data.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			// conn.Close()
			return
		}

		// sendData = strings.TrimSpace(sendData)
		// fmt.Printf("Sending: %s\n", sendData)
		_, err = conn.Write([]byte(sendData))
		if err != nil {
			fmt.Println("Error sending:", err)
			// conn.Close()
			return
		}
		fmt.Println("Message sent to client:", sendData)
		if sendData == "byebye\n" { //发送，其实上退出了，但是在Terminal上还没退出
			fmt.Println("Server said goodbye.")
			break
		}
		// 这里可以添加回应客户端的代码
	}
	conn.Close()
}
