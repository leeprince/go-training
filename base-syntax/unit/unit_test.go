/* 单元测试
1. 测试的文件名：{包名同名}"_test"
2. 测试方法：以 Test 开头，后面跟着要测试的方法名，并且后面跟着的测试方法名必须是可访问的方法（大写开头）
3. 运行方法：cd 到包名路径下，运行 $ go test [-v]
	可选参数：-v 输出运行详情
*/
/*
源码调试可使用包：delve

delve 工具：https://github.com/derekparker/delve

*/
package unit

import (
	"fmt"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	fmt.Println(GetUserInfo(1))
}

func TestGetUserName(t *testing.T) {
	fmt.Println(GetUserName(1))
}
