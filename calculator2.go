package main

import (
	"fmt"
	"math"
	"os"
)

//这个代码没做好。

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "operation")
		os.Exit(1)
	}
	str := os.Args[1]
	operation_before := false
	hasDecimalPoint := false
	var intPart1, intPart2, fracPart1, fracPart2 uint32
	var fracDigits1, fracDigits2 int
	var operation string
	validOperations := []string{"+", "-", "*", "/"}
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			if !operation_before && !hasDecimalPoint {
				intPart1 = intPart1*10 + uint32(str[i]-'0')
			} else if operation_before && !hasDecimalPoint {
				intPart2 = intPart2*10 + uint32(str[i]-'0')
			} else if !operation_before && hasDecimalPoint {
				fracPart1 = fracPart1*10 + uint32(str[i]-'0')
				fracDigits1++
			} else if operation_before && hasDecimalPoint {
				fracPart2 = fracPart2*10 + uint32(str[i]-'0')
				fracDigits2++
			}
		} else if contains(validOperations, string(str[i])) {
			operation = string(str[i])
			operation_before = true
			hasDecimalPoint = false
		} else if str[i] == '.' {
			hasDecimalPoint = true
		} else {
			fmt.Println("Invalid input")
			os.Exit(1)
		}
	}
	number1 := StrToNum(intPart1, fracPart1, fracDigits1, hasDecimalPoint)
	fmt.Println("Number1:", number1)
	number2 := StrToNum(intPart2, fracPart2, fracDigits2, hasDecimalPoint)
	fmt.Println("Number2:", number2)
	fmt.Println("Result:", calculate(number1, operation, number2))

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func StrToNum(intPart uint32, fracPart uint32, fracDigits int, hasDecimalPoint bool) interface{} {
	var result interface{}
	if hasDecimalPoint {
		fracPartF64 := float64(fracPart)
		result = float64(intPart) + fracPartF64/math.Pow10(fracDigits) // 使用 math.Pow10 计算整数次幂
	} else {
		if intPart > math.MaxInt32 {
			fmt.Println("Overflow: intPart exceeds the max value of int64")
			os.Exit(1)
		}
		result = int64(intPart)
	}
	return result
}

func calculate(number1 interface{}, operation string, number2 interface{}) interface{} {
	switch operation {
	case "+":
		return performOperation(number1, number2, func(a, b float64) float64 { return a + b })
	case "-":
		return performOperation(number1, number2, func(a, b float64) float64 { return a - b })
	case "*":
		return performOperation(number1, number2, func(a, b float64) float64 { return a * b })
	case "/":
		if number2 == "0" {
			fmt.Println("Division by zero")
			return nil
		}
		return performDivision(number1, number2)
	default:
		fmt.Println("Invalid operation:", operation)
		return nil
	}
}

// 使用函数组合来处理加减乘操作，确保转换为float64进行计算
func performOperation(a interface{}, b interface{}, operation func(float64, float64) float64) interface{} {
	floatA, floatB, ok := toFloat64(a, b)
	if !ok {
		return nil
	}
	return operation(floatA, floatB)
}

// 处理除法操作，需要额外检查除数是否为零
func performDivision(a interface{}, b interface{}) interface{} {
	floatA, floatB, ok := toFloat64(a, b)
	if !ok || floatB == 0.0 {
		return nil
	}
	return floatA / floatB
}

// 将两个接口转换为float64，如果都成功则返回true
func toFloat64(a interface{}, b interface{}) (float64, float64, bool) {
	floatA, okA := toFloat(a)
	floatB, okB := toFloat(b)
	if okA && okB {
		return floatA, floatB, true
	}
	return 0, 0, false
}

// 将接口转换为float64，如果可能则返回true
func toFloat(i interface{}) (float64, bool) {
	switch v := i.(type) {
	case int:
		return float64(v), true
	case int8:
		return float64(v), true
	case int16:
		return float64(v), true
	case int32:
		return float64(v), true
	case int64:
		return float64(v), true
	case float32:
		return float64(v), true
	case float64:
		return v, true
	default:
		return 0, false
	}
}
