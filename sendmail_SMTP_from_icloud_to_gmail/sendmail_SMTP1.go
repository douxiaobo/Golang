//OK

package main

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"os"
)

func main() {
	// 邮件服务器配置
	host := "smtp.mail.me.com"
	port := "587" // iCloud SMTP端口，使用TLS

	// 邮件发送者和接收者
	from := "dxb_1020@icloud.com"
	to := "douxiaobo@gmail.com"

	// 邮件内容
	subject := "Hello from Go"
	body := "Hello, this is a test email sent using Go's standard library!"

	// 构建邮件头部
	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = subject

	// 将邮件头部转换为字符串
	//邮件头部:
	for key, value := range header {
		body = key + ": " + value + "\r\n" + body
	}

	// 连接到SMTP服务器
	c, err := smtp.Dial(host + ":" + port)
	if err != nil {
		log.Fatal(err)
	}

	// 使用TLS加密
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 在生产环境中应设置为false
		ServerName:         host,
	}
	if err := c.StartTLS(tlsConfig); err != nil {
		log.Fatal(err)
	}

	// 认证
	auth := smtp.PlainAuth("", from, os.Getenv("SMTP_PASSWORD"), host)
	if err := c.Auth(auth); err != nil {
		log.Fatal(err)
	}

	// 发送邮件
	if err := c.Mail(from); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt(to); err != nil {
		log.Fatal(err)
	}

	w, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}

	// 写入邮件内容
	_, err = w.Write([]byte(body))
	if err != nil {
		log.Fatal(err)
	}

	// 结束邮件内容的写入
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}

	// 邮件发送完毕
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")
}
