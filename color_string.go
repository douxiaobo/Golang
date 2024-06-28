package main

import (
	"fmt"
	"strings"
)

func colorSet(str, selColor string) string {
	colorDic := map[string]string{
		"RED":     "\033[31m", // 红色
		"GREEN":   "\033[32m", // 绿色
		"YELLOW":  "\033[33m", // 黄色
		"BLUE":    "\033[34m", // 蓝色
		"FUCHSIA": "\033[35m", // 紫红色
		"CYAN":    "\033[36m", // 青蓝色
		"WHITE":   "\033[37m", // 白色
		"NORMAL":  "\033[0m",  // 终端默认颜色
	}

	selColor = strings.ToUpper(selColor)
	if color, ok := colorDic[selColor]; ok {
		return fmt.Sprintf("%s%s%s", color, str, colorDic["NORMAL"])
	} else {
		fmt.Println("没有找到对应颜色，采用终端默认颜色...")
		return fmt.Sprintf("%s%s%s", colorDic["NORMAL"], str, colorDic["NORMAL"])
	}
}

func main() {
	fmt.Println(colorSet("这一句话是红色", "RED"))
	fmt.Println(colorSet("这一句话是绿色", "green"))
	fmt.Println(colorSet("这一句话是黄色", "yellow"))
	fmt.Println(colorSet("这一句话是蓝色", "blue"))
	fmt.Println(colorSet("这一句话是紫红色", "fuchsia"))
	fmt.Println(colorSet("这一句话是颜色未设置", "test"))
}