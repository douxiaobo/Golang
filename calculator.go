package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("参数个数错误")
		return
	}
	fmt.Println("Number1= \"" + args[1] + "\" ,Operation= \"" + args[2] + "\" ,Number2= \"" + args[3] + "\"")
	result := calculate(args[1], args[2], args[3])
	fmt.Println("Result= \"" + result + "\"")
}

func calculate(number1 string, operation string, number2 string) string {
	num_1, err := strconv.ParseInt(number1, 10, 64)
	if err != nil {
		return "Error:" + err.Error()
	}
	num_2, err := strconv.ParseInt(number2, 10, 64)
	if err != nil {
		return "Error:" + err.Error()
	}
	switch operation {
	case "+":
		return fmt.Sprintf("%v", (num_1 + num_2))
	case "-":
		return fmt.Sprintf("%v", (num_1 - num_2))
	case "*":
		return fmt.Sprintf("%v", (num_1 * num_2))
	case "/":
		return fmt.Sprintf("%v", (num_1 / num_2))
	default:
		return "error"
	}
}
