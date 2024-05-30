package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {
	info := Website{
		Name:   "Golang Website",
		Url:    "https://golang.org",
		Course: []string{"Golang basics", "Web development with Golang", "Data Structures and Algorithms with Golang"}}
	file, err := os.Create("website.xml")
	if err != nil {
		fmt.Println("Error creating file:", err.Error())
		return
		panic(err)
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("Error encoding XML:", err.Error())
		return
		panic(err)
	} else {
		fmt.Println("XML file created successfully")
	}
}

// XML file created successfully
