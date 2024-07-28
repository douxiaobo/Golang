package main

type Website struct {
	Language    string
	Title       string
	Menu        []MenuLink
	Content     string
	ContentName string
	FooterLinks []FooterLink
	CookieName  string
	Travel      []TravelEntry
}

type WorkContent struct {
	Zh string `yaml:"zh"`
	En string `yaml:"en"`
	Es string `yaml:"es"`
}

type TravelEntry struct {
	Id      string `json:"Id"`
	Date    string `json:"Date"`
	Country string `json:"Country"`
}

type TravelData struct {
	Zh []TravelEntry `json:"zh"`
	En []TravelEntry `json:"en"`
	Es []TravelEntry `json:"es"`
}

type MenuLink struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type FooterLink struct {
	Code string `json:"code"` // 假设我们使用"code"作为JSON中的键，代表语言代码
	Name string `json:"name"` // 假设我们使用"name"作为JSON中的键，代表语言名称
}

type FooterLinks []FooterLink

type MarathonData struct {
	Zh []MarathonEntry `json:"zh"`
	En []MarathonEntry `json:"en"`
	Es []MarathonEntry `json:"es"`
}

type MarathonEntry struct {
	Id       string `json:"Id"`
	Date     string `json:"Date"`
	City     string `json:"City"`
	Marathon string `json:"Marathon"`
	Project  string `json:"Project"`
}

var user Website

type LanguageMap map[string]string

var langrange = [...]string{"en", "zh", "es"}

var containtlist = [...]string{"home", "about", "work", "travel", "music", "programming", "school", "sport"}
