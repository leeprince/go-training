package unit

import (
	"encoding/json"
	"fmt"
)

type s struct {
	Name       string
	Age        int    `json:"myAge"`
	Mage       int    `json:"MMyage"`
	nameIgnore string `json:"NameIgnore"`
	Ignore     string `json:"-"`
}

func Json() {
	/*
		结构体转json
			结构体的字段必须是可访问的（大写开头），否则无法转换而忽略
	*/
	// sj := s{Name: "prince", Age: 16, Mage: 20, nameIgnore: "ig", Ignore: "mIg"} // 必须是：nameIgnore，NameIgnore是错误的
	sj := s{"prince", 16, 20, "ig", "mIg"}
	sToj, _ := json.Marshal(sj)
	fmt.Println("结构体转json:", string(sToj))

	/*
		json转结构体
			结构体的字段必须是可访问的（大写开头），否则无法转换而忽略
			不可访问的字段使用可访问方式的匿名字段（大写开头）同样不可json转结构体
	*/
	// 不可访问的 nameIgnore 使用匿名字段：NameIgnore 同样不可json转结构体
	/*js := []byte(`
		{"Name": "prince", "myAge": 18, "MMyage": 20, "NameIgnore": "ig", "Ignore": "mIg"}
	`)*/
	// 正常转换
	/*js := []byte(`
		{"Name": "prince", "myAge": 18, "MMyage": 20}
	`)*/
	// 正常转换
	// json转结构体 忽略可访问字段前提下的大小写
	js := []byte(`
		{"Name": "prince", "MyAge": 18, "mMyage": 20}
	`)
	ss := s{}
	err := json.Unmarshal(js, &ss)
	if err != nil {
		fmt.Println("json转结构体失败")
	}
	fmt.Println("json转结构体:", ss)

}
