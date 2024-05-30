package main

import (
	"fmt"
	"net"
	"time"
)

func echo(conn *net.TCPConn) {
	tick := time.Tick(5 * time.Second)
	for now := range tick {
		n, err := conn.Write([]byte(now.String()))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
		fmt.Println("Wrote", n, "bytes to client")
		fmt.Println("send %d bytes to %s\n", n, conn.RemoteAddr())
	}
}
func main() {
	address := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8000,
	}
	listener, err := net.ListenTCP("tcp4", &address)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())
		go echo(conn)
	}
}

// go run server.go
// Accepted connection from 127.0.0.1:64305
// Wrote 52 bytes to client
// send %d bytes to %s
//  52 127.0.0.1:64305
// Wrote 51 bytes to client
// send %d bytes to %s
//  51 127.0.0.1:64305
// Wrote 52 bytes to client
// send %d bytes to %s
//  52 127.0.0.1:64305
// Wrote 51 bytes to client
// send %d bytes to %s
//  51 127.0.0.1:64305
// Wrote 52 bytes to client
// send %d bytes to %s
//  52 127.0.0.1:64305
// Wrote 52 bytes to client
// send %d bytes to %s
//  52 127.0.0.1:64305
// Error writing to client: write tcp4 127.0.0.1:8000->127.0.0.1:64305: write: broken pipe
