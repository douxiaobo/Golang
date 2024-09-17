package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Location struct {
	City    string
	Country string
}

type PageData struct {
	Location string
	Message  string
}

// var tmpl = template.Must(template.ParseFiles("templates/index.html"))
var tmpl = template.Must(template.ParseFiles("./index.html"))

func getLocationByIP(ip string) (Location, error) {
	// 这里应该调用一个实际的地理定位API，例如ip-api.com
	// 由于实际API可能需要认证或其他配置，此处仅作示例
	return Location{"Unknown City", "Unknown Country"}, nil
}

func getLocationByJS() string {
	// JavaScript 会在这里执行，获取地理位置
	return `
	<script>
		function getLocation() {
			if (navigator.geolocation) {
				navigator.geolocation.getCurrentPosition(showPosition);
			} else { 
				document.getElementById('location').innerHTML = "Geolocation is not supported by this browser.";
			}
		}

		function showPosition(position) {
			fetch('/getLocation', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify({
					lat: position.coords.latitude,
					lng: position.coords.longitude
				})
			}).then(response => response.json())
			.then(data => document.getElementById('location').innerHTML = data.location);
		}
	</script>
	`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// 获取客户端IP
	clientIP := r.RemoteAddr

	// 尝试从IP获取位置
	loc, _ := getLocationByIP(clientIP)

	// 如果有位置功能，使用JavaScript获取位置
	jsCode := getLocationByJS()

	data := PageData{
		Location: fmt.Sprintf("%s, %s", loc.City, loc.Country),
		Message:  "你好",
	}

	tmpl.Execute(w, data)
	fmt.Fprint(w, jsCode)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Starting HTTP server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
