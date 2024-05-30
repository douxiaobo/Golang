package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}
	service := strings.Split(os.Args[1], ":")
	tcpAdd, err := net.ResolveTCPAddr("tcp4", service[0]+":"+service[1])
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAdd)
	checkError(err)
	defer conn.Close()
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}

// go run client.go 127.0.0.1:8000
// ^Csignal: interrupt
