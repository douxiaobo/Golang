package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.ListenUDP("udp4", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)

	}
}
func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[:])
	if err != nil {
		fmt.Println("Error reading from UDP connection:", err)
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
