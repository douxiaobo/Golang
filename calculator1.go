package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<expression>")
		return
	}
	str := os.Args[1]
	fmt.Println(str)
	operation_before := false
	number1 := 0
	number2 := 0
	var operation string                            // 添加类型声明
	validOperations := []string{"+", "-", "*", "/"} // 添加有效运算符列表
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			if operation_before {
				number2 = number2*10 + int(str[i]-'0')
			} else {
				number1 = number1*10 + int(str[i]-'0')
			}
		} else if contains(validOperations, string(str[i])) { // 检查运算符是否有效
			operation_before = true
			operation = string(str[i])
		} else {
			fmt.Println("Invalid character:", str[i])
			return // 处理无效运算符的情况
		}
	}
	fmt.Println("Result:", calculate(number1, operation, number2))
}

// 定义一个辅助函数来检查字符串是否存在于切片中
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func calculate(number1 int, operation string, number2 int) int {
	switch operation {
	case "+":
		return number1 + number2
	case "-":
		return number1 - number2
	case "*":
		return number1 * number2
	case "/":
		if number2 == 0 {
			fmt.Println("Division by zero")
			return 0
		}
		return number1 / number2
	default:
		fmt.Println("Invalid operation:", operation)
		return 0
	}
}
