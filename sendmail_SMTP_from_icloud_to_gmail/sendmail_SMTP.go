// 写代码成功了。

package main

import (
	"bytes"
	"crypto/tls"
	"log"
	"net/smtp"
	"os"
	"text/template"
)

func main() {
	// 邮件内容模板
	const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}
`
	// 邮件数据
	data := struct {
		From    string
		To      string
		Subject string
		Body    string
	}{
		From:    "dxb_1020@icloud.com",
		To:      "douxiaobo@gmail.com",
		Subject: "Hello, Go SMTP!",
		Body:    "Hello, this is a test email sent using Go's standard library!",
	}

	// 创建邮件正文
	t, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		log.Fatal(err)
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		log.Fatal(err)
	}

	// 发送邮件
	auth := smtp.PlainAuth("", data.From, os.Getenv("SMTP_PASSWORD"), "smtp.mail.me.com")

	// 使用SMTP服务器发送邮件
	smtpHost := "smtp.mail.me.com"
	smtpPort := "587" // 使用587端口进行STARTTLS加密

	// 建立SMTP连接
	c, err := smtp.Dial(smtpHost + ":" + smtpPort)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// 升级连接到TLS
	if err = c.StartTLS(&tls.Config{ServerName: smtpHost}); err != nil {
		log.Fatal(err)
	}

	// 进行认证
	if err := c.Auth(auth); err != nil {
		log.Fatal(err)
	}

	// 设置邮件发送者和接收者
	if err := c.Mail(data.From); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt(data.To); err != nil {
		log.Fatal(err)
	}

	// 创建写入邮件内容的Writer
	w, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}

	// 写入邮件内容
	_, err = w.Write(tpl.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	// 结束邮件内容的写入
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 邮件发送完毕
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")
}

// 通过在终端中运行以下命令来设置环境变量：
// export SMTP_PASSWORD=your_app_specific_password
