package main

import (
	"fmt"
	"strings"
)

// passwdPre 形成密码前缀，通过对密码中每个字符进行转换得到一个新字符串
func passwdPre(pwd string) string {
	var vret strings.Builder
	for _, char := range pwd {
		switch {
		case char == 'a' || char == 'b' || char == 'c':
			vret.WriteRune('!')
		case char == 'd' || char == 'e' || char == 'f':
			vret.WriteRune('@')
		case char == 'g' || char == 'h' || char == 'i':
			vret.WriteRune('#')
		case char == 'j' || char == 'k' || char == 'l':
			vret.WriteRune('%')
		case char == 'm' || char == 'n' || char == 'o':
			vret.WriteRune('^')
		case char == 'p' || char == 'q' || char == 'r':
			vret.WriteRune('&')
		case char == 's' || char == 't' || char == 'u':
			vret.WriteRune('*')
		case char == 'v' || char == 'w' || char == 'x':
			vret.WriteRune('>')
		case char == 'y' || char == 'z':
			vret.WriteRune('?')
		case char >= 'A' && char <= 'Z':
			vret.WriteRune(rune(int(char) - 'A' + 1 + 'a'))
		case char == 'Z':
			vret.WriteRune('a')
		default:
			vret.WriteRune(char)
		}
	}
	return vret.String()
}

// changeTxt 根据传入的密码和两个字符串，对密码中字符进行转换
func changeTxt(pwd, str1, str2 string) string {
	var vret strings.Builder
	for _, char := range pwd {
		if i := strings.IndexRune(str1, char); i >= 0 {
			vret.WriteRune(rune(str2[i]))
		} else {
			vret.WriteRune(char)
		}
	}
	return vret.String()
}

// changePassword 加密程序
func changePassword(pwd string) string {
	if pwd == "" {
		return "-1"
	}
	vpre := passwdPre(pwd)
	vlen := len(pwd)
	vret := vpre

	converters := []struct {
		str1, str2 string
	}{
		{"1234567890abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz1234567890"},
		{"1234567890abcdefghijklmnopqrstuvwxyz", "qwertyuiopasdfghjklmnbvcxz0987654321"},
		{"1234567890abcdefghijklmnopqrstuvwxyz", "1qaz2wsx3edc4rfv5tgb6yhn7ujm8ik9ol0p"},
		{"1234567890abcdefghijklmnopqrstuvwxyz", "pl0okm9ijn8uhb7ygv6tfc5rdx4esz3wa2q1"},
	}

	for _, converter := range converters {
		vstr := changeTxt(pwd, converter.str1, converter.str2)
		if vlen <= 4 {
			vret += vstr[:vlen]
		} else {
			vret += vstr[:4]
		}
	}

	return vret
}

func main() {
	for {
		var pwd string
		fmt.Print("请录入密码：")
		fmt.Scanln(&pwd)
		if pwd == "q" {
			fmt.Println("退出程序...")
			break
		} else {
			pwdnew := changePassword(pwd)
			fmt.Printf("您录入的密码是: %s, 该密码加密后为：%s\n", pwd, pwdnew)
		}
	}
}

// douxiaobo@192 Golang % go run password_change.go
// 请录入密码：testT123@/
// 您录入的密码是: testT123@/, 该密码加密后为：*@**u123@/4o347g87if8iz7sz
// 请录入密码：ttT22#liK
// 您录入的密码是: ttT22#liK, 该密码加密后为：**u22#%#l44Tb77TwiiTqzzTl
// 请录入密码：q
// 退出程序...
// douxiaobo@192 Golang %

// 错误
