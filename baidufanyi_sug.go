package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

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

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return
	}

	for _, item := range result["data"].([]interface{}) {
		itemMap := item.(map[string]interface{})
		v := itemMap["v"].(string)
		fmt.Println(v)
	}

	fmt.Println()

	// 定义一个结构体来解析JSON响应
	var response struct {
		Errno int `json:"errno"`
		Data  []struct {
			K string `json:"k"`
			V string `json:"v"`
		} `json:"data"`
		Logid int `json:"logid"`
	}

	// 解析JSON
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 遍历Data中的每个条目，并转换V字段中的Unicode转义序列为字符串
	for _, item := range response.Data {
		decodedV, err := strconv.Unquote(`"` + item.V + `"`)
		if err != nil {
			fmt.Println("Error decoding string:", err)
			continue
		}
		fmt.Printf("%s: %s\n", item.K, decodedV)
	}
}
