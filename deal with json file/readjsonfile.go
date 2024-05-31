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
	info := []Website{
		{Name: "Google", Url: "https://www.google.com", Course: []string{"Go", "Python"}},
		{Name: "Youtube", Url: "https://www.youtube.com", Course: []string{"Java", "C++"}},
		{Name: "Facebook", Url: "https://www.facebook.com", Course: []string{"C", "C++"}},
		{Name: "Twitter", Url: "https://www.twitter.com", Course: []string{"C", "C++"}},
		{Name: "Instagram", Url: "https://www.instagram.com", Course: []string{"C", "C++"}},
		{Name: "Linkedin", Url: "https://www.linkedin.com", Course: []string{"C", "C++"}},
		{Name: "Github", Url: "https://www.github.com", Course: []string{"C", "C++"}},
		{Name: "Stackoverflow", Url: "https://www.stackoverflow.com", Course: []string{"C", "C++"}},
	}
	filePtr, err := os.Create("website.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filePtr.Close()
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success.")
	}

}

// success.
