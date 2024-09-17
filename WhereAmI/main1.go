package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

// 位置结构体
type Location struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	IP        string  `json:"ip,omitempty"`
	Country   string  `json:"country,omitempty"`
}

// 获取客户端 IP
func getClientIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	return ip
}

// 模拟从 IP 获取位置
func getLocationFromIP(ip string) (*Location, error) {
	// 在这里可以调用第三方 API 获取位置信息
	// 这是一个示例，实际应用中需要替换为有效的 API 调用
	return &Location{
		IP:      ip,
		Country: "中国", // 示例中使用静态值
	}, nil
}

// 处理请求
func locationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 检查请求参数以区分两种设备
	positionType := r.URL.Query().Get("type") // 设备类型：位置功能或无位置功能

	var loc Location
	var err error

	if positionType == "gps" {
		// 有位置功能的设备，通常会通过前端 JavaScript 获取 GPS 位置信息并发送
		// 这里简化处理，直接假设 GPS 信息
		loc = Location{
			Latitude:  39.9042,  // 示例纬度
			Longitude: 116.4074, // 示例经度
		}
	} else {
		// 没有位置功能的设备，从 IP 获取位置
		ip := getClientIP(r)
		loc, err = getLocationFromIP(ip)
		if err != nil {
			http.Error(w, "无法获取位置", http.StatusInternalServerError)
			return
		}
	}

	// 将位置数据返回给客户端
	json.NewEncoder(w).Encode(loc)
}

func main() {
	http.HandleFunc("/location", locationHandler)
	fmt.Println("服务器正在监听 :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

//测试接口：在浏览器中访问 http://localhost:8080/location?type=gps 来测试 GPS 设备，或访问 http://localhost:8080/location?type=ip 来模拟无位置功能的设备（从 IP 获取位置）。
