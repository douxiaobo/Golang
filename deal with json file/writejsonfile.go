package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {
	filePtr, err := os.Open("./website.json")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer filePtr.Close()

	var info []Website
	decode := json.NewDecoder(filePtr)
	err = decode.Decode(&info)
	if err != nil {
		fmt.Println("Error decoding JSON")
		return
	} else {
		fmt.Println("JSON data is success.")
		fmt.Println(info)
	}
}

// go run writejsonfile.go
// JSON data is success.
// [{Google https://www.google.com [Go Python]} {Youtube https://www.youtube.com [Java C++]} {Facebook https://www.facebook.com [C C++]} {Twitter https://www.twitter.com [C C++]} {Instagram https://www.instagram.com [C C++]} {Linkedin https://www.linkedin.com [C C++]} {Github https://www.github.com [C C++]} {Stackoverflow https://www.stackoverflow.com [C C++]}]
