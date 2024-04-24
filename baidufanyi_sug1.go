package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	Errno int         `json:"errno"`
	Data  []DataEntry `json:"data"`
	Logid int64       `json:"logid"`
}

type DataEntry struct {
	K string `json:"k"`
	V string `json:"v"`
}

func main() {

	url := "https://fanyi.baidu.com/sug"
	method := "POST"

	payload := strings.NewReader("kw=Hello")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "BAIDUID=4C26A499B77432870FC4DC861A9DEA01:FG=1; BAIDUID_BFESS=4C26A499B77432870FC4DC861A9DEA01:FG=1")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Output the decoded JSON data
	jsonData, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(jsonData))
}
