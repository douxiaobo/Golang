package main

import (
	"fmt"
	"unicode"
	"strings"
)


func check_len(password string) bool{
	if len(password) < 8{
		// fmt.Println("密码长度不能小于8位");
		return false;
	} else {
		// fmt.Println("密码长度符合要求");
		return true;
	}
}

func check(password string) bool{
	var check=[4]int {0,0,0,0};
	for _,ch:=range password {
		if unicode.IsLower(ch){
			check[0]=1;
		} 
		if unicode.IsUpper(ch){
			check[1]=1;
		} 
		if unicode.IsDigit(ch){
			check[2]=1;
		} 
		if !(unicode.IsLetter(ch)||unicode.IsDigit(ch)||unicode.IsSpace(ch)) {
			check[3]=1;
		}
	}
	var sum int=0;
	for i:=0;i<4;i++{
		sum+=check[i];
	}
	if sum<4{
		return false;
	} else {
		return true;
	}
}

func check_rep(password string)bool{
	for i:=0;i<len(password)-4;i++{
		var str1=password[i:i+4];
		var str2=password[i+4:len(password)];
		if strings.Contains(str2,str1){
			return false;
		}
	}
	return true;
}

func main(){
	var msg string=`请设置密码，密码要求符合以下条件：
    1、密码长度不小于8位
    2、密码必须由字母大、小写、数字、其它符号组成
    3、密码中不能重复包含长度超4的子串`;
	fmt.Println(msg);
	for true {
		fmt.Print("请输入密码：");
		var pwd string;
		fmt.Scanln(&pwd);
		if pwd=="q"{
			fmt.Println("退出程序...");
			break;
		}
		if !check_len(pwd){
			fmt.Println("密码长度不够8位！请重新录入.");
			continue;
		}
		if !check(pwd){
			fmt.Println("密码必须由字母大、小写、数字、其它符号组成！请重新录入.");
			continue;
		}
		if !check_rep(pwd){
			fmt.Println("密码包含两个以上重复子串（4位以上的子串）！请查看并重新录入重新录入.");
			continue;
		}
		fmt.Println("密码正确。");
		break;
	}
}