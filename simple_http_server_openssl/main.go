package main

import (
	"log"
	"net/http"
)

func main() {
	//启动服务器
	srv := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(handler)}
	//用TSL启动服务器
	log.Printf("Starting server on https://localhost:8080")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)

	w.Write([]byte("Hello, world!"))
}

// Terminal in Mac OS:
// openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
// ....+.........+..+.........+.+...+...+........+.+.....+.+..+.+......+......+.....+....+.....+......+++++++++++++++++++++++++++++++++++++++*..+.+.....+.+...+...+..+.+.....+.+...........+.........+++++++++++++++++++++++++++++++++++++++*.............+.........+.....+..........+......+...+...........+...+............+...+.+....................+......+.+.....+..........+.........+...+..+....+......+......+..+..........+..+.+..+.............+..+.+............+..+......+....+...+..............+...................+..............+.......+...............+......+..+.+..+........................+.+...+..+............+.......+........+...............++++++
// .....+++++++++++++++++++++++++++++++++++++++*...+.......+...+..+.........+...+.+............+..+...+.........+.+...+............+..+...+..................+....+..+.+.....+.......+++++++++++++++++++++++++++++++++++++++*..+............................+.....+............+.+......+...++++++
// -----
// You are about to be asked to enter information that will be incorporated
// into your certificate request.
// What you are about to enter is what is called a Distinguished Name or a DN.
// There are quite a few fields but you can leave some blank
// For some fields there will be a default value,
// If you enter '.', the field will be left blank.
// -----
// Country Name (2 letter code) [AU]:CN
// State or Province Name (full name) [Some-State]:Zhejiang
// Locality Name (eg, city) []:Ningbo
// Organization Name (eg, company) [Internet Widgits Pty Ltd]:Microsoft
// Organizational Unit Name (eg, section) []:MW
// Common Name (e.g. server FQDN or YOUR name) []:DXB
// Email Address []:dxb_1020@icloud.com

// http://0.0.0.0:8080/
// Client sent an HTTP request to an HTTPS server.

// https://localhost:8080/
// Hello, world!

// https://0.0.0.0:8080/
// Hello, world!

// go run main.go
// 2024/05/18 16:26:13 Starting server on https://localhost:8080
// 2024/05/18 16:26:40 http: TLS handshake error from [::1]:56450: remote error: tls: unknown certificate
// 2024/05/18 16:26:40 http: TLS handshake error from [::1]:56451: remote error: tls: unknown certificate
// 2024/05/18 16:27:04 http: TLS handshake error from [::1]:56459: remote error: tls: unknown certificate
// 2024/05/18 16:27:04 http: TLS handshake error from [::1]:56460: remote error: tls: unknown certificate
// 2024/05/18 16:27:12 http: TLS handshake error from [::1]:56465: remote error: tls: unknown certificate
// 2024/05/18 16:27:12 http: TLS handshake error from [::1]:56466: remote error: tls: unknown certificate
// 2024/05/18 16:27:12 Got connection: HTTP/2.0
// 2024/05/18 16:27:12 Got connection: HTTP/2.0
// 2024/05/18 16:27:28 http: TLS handshake error from [::1]:56444: EOF
// 2024/05/18 16:28:15 Got connection: HTTP/2.0
// 2024/05/18 16:28:16 Got connection: HTTP/2.0
// 2024/05/18 16:28:33 http: TLS handshake error from 127.0.0.1:56475: EOF
// 2024/05/18 16:30:15 http: TLS handshake error from 127.0.0.1:56535: remote error: tls: unknown certificate
// 2024/05/18 16:30:15 http: TLS handshake error from 127.0.0.1:56536: remote error: tls: unknown certificate
// 2024/05/18 16:30:21 http: TLS handshake error from 127.0.0.1:56537: remote error: tls: unknown certificate
// 2024/05/18 16:30:21 http: TLS handshake error from 127.0.0.1:56538: remote error: tls: unknown certificate
// 2024/05/18 16:30:21 Got connection: HTTP/2.0
// 2024/05/18 16:30:21 Got connection: HTTP/2.0
