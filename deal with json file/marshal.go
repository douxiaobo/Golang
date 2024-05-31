package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{"Server1", "192.168.1.1"})
	s.Servers = append(s.Servers, Server{"Server2", "192.168.1.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

// {"Servers":[{"ServerName":"Server1","ServerIP":"192.168.1.1"},{"ServerName":"Server2","ServerIP":"192.168.1.2"}]}
