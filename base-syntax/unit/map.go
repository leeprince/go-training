package unit

import "fmt"

type mapData struct {
	m map[int]string
}
/*
Go语言中的map是引用类型，必须初始化才能使用
 */
func Map() {
	/*var goClassMap map[string]string // 集合
	goClassMap = make(map[string]string)*/
	goClassMap := make(map[string]string) // 集合【推荐】
	// 添加数据
	goClassMap["name"] = "prince"
	goClassMap["like"] = "go"
	fmt.Println(goClassMap)
	// 输出
	for index, value := range goClassMap {
	fmt.Printf("index : %v, value : %v \n", index, value) }

	/*
	检查键是否存在
	 */
	userInfo := map[string]string{
		"name": "prince",
		"like": "golang",
	}
	checkKey, ok := userInfo["name001"]
	if ok {
		fmt.Println("键name001存在, 键值为：", checkKey)
	} else {
		fmt.Println("键name001存在")
	}
	checkKeyName, ok := userInfo["name"]
	if ok {
		fmt.Println("键name不存在, 键值为：", checkKeyName)
	} else {
		fmt.Println("键name存在")
	}

	data := map[int]string{0:"prince", 1:"golang"}
	mapChange(data)
	fmt.Println("修改map值后的结果：", data)
	
	mapStruct := mapData{m: map[int]string{0:"prince", 1:"golang"}}
	mapStructChange(mapStruct)
	fmt.Println("修改map值后的结果：", mapStruct)
	
	mapStructQuote := mapData{m: map[int]string{0:"prince", 1:"golang"}}
	mapStructChangeQuote(&mapStructQuote)
	fmt.Println("修改map值后的结果：", mapStruct)
}

func mapChange(data map[int]string)  {
	data[0] = "100"
}

func mapStructChange(data mapData)  {
	s := data.m
	fmt.Println("s:", s)
	s1 := data.m
	s1[0] = "100"
	fmt.Println("s:", s1)

	data.m[0] = "101"
}

func mapStructChangeQuote(data *mapData)  {
	s := data.m
	fmt.Println("s:", s)
	s1 := data.m
	s1[0] = "100"
	fmt.Println("s:", s1)

	data.m[0] = "101"
}
