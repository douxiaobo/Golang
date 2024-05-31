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
	str := `{
		"servers":
		[
			{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
			{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"},
			{"ServerName":"Server1","ServerIP":"192.168.1.1"},
			{"ServerName":"Server2","ServerIP":"192.168.1.2"},
			{"ServerName":"Server3","ServerIP":"192.168.1.3"}
		]
	}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

// {[{Shanghai_VPN 127.0.0.1} {Beijing_VPN 127.0.0.2} {Server1 192.168.1.1} {Server2 192.168.1.2} {Server3 192.168.1.3}]}
