package main

import "testing"

/*
main 包名的测试方法统一
	func TestMain(m *testing.M) {
		...
	}
 */
func TestMain(m *testing.M) {
	TestInt()
}