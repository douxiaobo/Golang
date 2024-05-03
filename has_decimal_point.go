package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s <number>\n", args[0])
		os.Exit(1)
	}

	var intPart, fracPart uint32
	var fracDigits int
	hasDecimalPoint := false

	for _, c := range strings.TrimSpace(args[1]) {
		if unicode.IsDigit(c) && !hasDecimalPoint { // 整数
			intPart = intPart*10 + uint32(c-'0')
		} else if unicode.IsDigit(c) && hasDecimalPoint { // 小数点
			fracPart = fracPart*10 + uint32(c-'0')
			fracDigits++
		} else if c == '.' {
			hasDecimalPoint = true
			// fmt.Println("has decimal point")
		} else {
			fmt.Println("error")
			return
		}
	}

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

	// 分别处理打印逻辑
	switch v := result.(type) {
	case float64:
		fmt.Printf("num: %.6f\n", v) // 控制浮点数输出的精度为6位小数
	case int64:
		fmt.Println("num:", v)
	default:
		fmt.Println("Unexpected result type")
	}

}

// 主要变化如下：

// 1. 导入必要的 Go 语言库：fmt, os, strconv, strings, unicode, 和 math。

// 2. 使用 os.Args 获取命令行参数，替代 Rust 中的 std::env::args()。

// 3. 改用 Go 语言的条件判断语句和循环结构。

// 4. 使用 unicode.IsDigit 函数检查字符是否为数字。

// 5. 计算整数部分时，将字符转换为数字的方式调整为 c - '0'，这是 Go 语言中常见的做法。

// 6. 使用 math.Pow10 替代 Rust 中的 10f32.powi(fracDigits)，实现相同的功能（计算以 10 为底的整数次幂）。

// 7. 最后，使用 fmt.Printf 打印结果，其中 %.6f 控制浮点数输出的精度为 6 位小数。根据实际需求，您可以调整这个数值或直接使用 %f。
