package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 创建一个HTTP请求
	req, err := http.NewRequest("GET", "http://www.baidu.com/", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 添加User-Agent头部（这里你可以指定一个真实的User-Agent字符串）
	// 如果你想模拟一个空的User-Agent，这通常不是一个好主意，因为它可能会被服务器拒绝
	// 但为了与Python代码保持一致，这里我们仍然留空
	req.Header.Set("User-Agent", "")

	// 发送请求并获取响应
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 打印状态码和原因
	fmt.Printf("Status: %d %s\n", resp.StatusCode, resp.Status)

	// 打印头部信息
	for k, values := range resp.Header {
		for _, v := range values {
			fmt.Printf("%s: %s\n", k, v)
		}
	}

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// fmt.Println(string(body))
}

// Status: 200 200 OK
// Traceid: 1715856019053282151414132690136321678603
// Connection: keep-alive
// Date: Thu, 16 May 2024 10:40:19 GMT
// Server: BWS/1.1
// Set-Cookie: BAIDUID=553DC12C89A75A214B10E7F1A39B6D5B:FG=1; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
// Set-Cookie: BIDUPSID=553DC12C89A75A214B10E7F1A39B6D5B; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
// Set-Cookie: PSTM=1715856019; expires=Thu, 31-Dec-37 23:55:55 GMT; max-age=2147483647; path=/; domain=.baidu.com
// Set-Cookie: BAIDUID=553DC12C89A75A21BDEE0ADEA1CA4994:FG=1; max-age=31536000; expires=Fri, 16-May-25 10:40:19 GMT; domain=.baidu.com; path=/; version=1; comment=bd
// X-Ua-Compatible: IE=Edge,chrome=1
// Content-Security-Policy: frame-ancestors 'self' https://chat.baidu.com http://mirror-chat.baidu.com https://fj-chat.baidu.com https://hba-chat.baidu.com https://hbe-chat.baidu.com https://njjs-chat.baidu.com https://nj-chat.baidu.com https://hna-chat.baidu.com https://hnb-chat.baidu.com http://debug.baidu-int.com;
// Content-Type: text/html; charset=utf-8
// X-Xss-Protection: 1;mode=block
// P3p: CP=" OTI DSP COR IVA OUR IND COM "
// P3p: CP=" OTI DSP COR IVA OUR IND COM "
