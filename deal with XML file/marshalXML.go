package main

import (
	"encoding/xml"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}
type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{ServerName: "server1", ServerIP: "192.168.1.1"})
	v.Svs = append(v.Svs, server{ServerName: "server2", ServerIP: "192.168.1.2"})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		panic(err)
	}
	// println(string(output))
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

// <servers version="1">
// <server>
// 	<serverName>server1</serverName>
// 	<serverIP>192.168.1.1</serverIP>
// </server>
// <server>
// 	<serverName>server2</serverName>
// 	<serverIP>192.168.1.2</serverIP>
// </server>
// </servers>
// <?xml version="1.0" encoding="UTF-8"?>
// <servers version="1">
// <server>
// 	<serverName>server1</serverName>
// 	<serverIP>192.168.1.1</serverIP>
// </server>
// <server>
// 	<serverName>server2</serverName>
// 	<serverIP>192.168.1.2</serverIP>
// </server>
// </servers>%
