package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/huawen0327/golee"
)

// Settings 配置
type Settings struct {
	Port   int
	Static string
	Root   string
}

func explainJSON() *Settings {
	var settings Settings

	// 解析json文件
	f, err := ioutil.ReadFile("settings.json")
	if err != nil {
		fmt.Println("配置文件不存在:", err)
	}

	// 去除json的注释
	re, err := regexp.Compile("//.*?\n")
	if err != nil {
		fmt.Println("正则匹配出错:", err)
	}
	allString := re.FindAll([]byte(f), -1)
	fString := string(f)
	for _, value := range allString {
		fString = strings.Replace(fString, string(value), "\n", 1)
	}

	// 解析Json
	err = json.Unmarshal([]byte(fString), &settings)
	if err != nil {
		fmt.Println("配置文件出错:", err)
	}

	return &settings
}

func main() {
	settings := explainJSON()
	g := golee.New()
	g.Group.Static(settings.Static, settings.Root)
	url(g.Router)
	urlGroup(g.Group)
	addr := ":" + strconv.Itoa(settings.Port)
	fmt.Println("address:" + strconv.Itoa(settings.Port))
	err := g.Run(addr)
	fmt.Println(err)
}
