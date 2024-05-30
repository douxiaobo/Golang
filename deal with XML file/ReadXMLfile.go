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
	file, err := os.Open("website.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	var website Website
	err = decoder.Decode(&website)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Successfully read XML file")
		fmt.Println(website)
	}

	fmt.Println(website.Name)
	fmt.Println(website.Url)
	fmt.Println(website.Course)
}

// Successfully read XML file
// {Golang Website https://golang.org [Golang basics Web development with Golang Data Structures and Algorithms with Golang]}
// Golang Website
// https://golang.org
// [Golang basics Web development with Golang Data Structures and Algorithms with Golang]
