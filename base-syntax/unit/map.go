package unit

import "fmt"

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


}
