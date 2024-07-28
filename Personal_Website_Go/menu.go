// menu.go

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Menus struct {
	Zh []MenuLink `json:"zh"`
	En []MenuLink `json:"en"`
	Es []MenuLink `json:"es"`
}

func readAndParseMenus() (Menus, error) {
	menuFile, err := os.Open("./public/json/Menu.json")
	if err != nil {
		return Menus{}, fmt.Errorf("error opening file: %w", err)
	}
	defer menuFile.Close()

	menuData, err := io.ReadAll(menuFile)
	if err != nil {
		return Menus{}, fmt.Errorf("error reading file: %w", err)
	}

	type TempMenu struct {
		Zh []MenuLink `json:"zh"`
		En []MenuLink `json:"en"`
		Es []MenuLink `json:"es"`
	}

	// 将JSON数据反序列化到临时结构体
	var tempMenus TempMenu
	err = json.Unmarshal(menuData, &tempMenus)
	if err != nil {
		return Menus{}, fmt.Errorf("解析JSON出错: %w", err)
	}

	// 将临时结构体转换为Menus结构体
	menus := Menus{
		Zh: tempMenus.Zh,
		En: tempMenus.En,
		Es: tempMenus.Es,
	}

	return menus, nil
}

// 获取菜单
func getMenu(menus Menus, language string) []MenuLink {
	var menuItems []MenuLink
	switch language {
	case "zh":
		menuItems = menus.Zh
	case "en":
		menuItems = menus.En
	case "es":
		menuItems = menus.Es
	default:
		menuItems = []MenuLink{}
	}
	return menuItems
}
