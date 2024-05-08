package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}
	port := 12345
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	for {
		//发送
		fmt.Print("Enter message to send to server (type 'byebye' to exit):")
		send_data := bufio.NewReader(os.Stdin)
		sendData, err := send_data.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			return
		}
		conn.Write([]byte(sendData))
		fmt.Print("Message sent to server:", sendData)
		if sendData == "byebye\n" { //正常
			fmt.Println("Cient said goodbye.")
			break
		}
		// reader := bufio.NewReader(conn)
		// msg, err := reader.ReadString('\n')
		// if err != nil {
		// 	fmt.Println("Error reading from server:", err)
		// 	return

		// }
		// fmt.Print(msg)
		//接收
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading from server:", err)
			return
		}
		recvData := string(buffer[:n])
		fmt.Println("Message received from server:", recvData)
		if recvData == "byebye\n" { //正常
			fmt.Println("Server closed the connection.")
			break
		}
	}
}
