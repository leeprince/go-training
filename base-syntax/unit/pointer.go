package unit

import "fmt"

func Pointer() {
	/* & *
	1. 指针与引用：&:返回变量的指针（内存地址） *：返回指针的原值（指针指向的内容）
	2. 变量的指针的类型（变量的指针类型表示为：*变量名） 等于 *数据类型（数据类型的指针类型表示为：*数据类型）
	*/
	a := 10 // a  0xc0000ae008 =>> 10
	b := &a // b  0xc0000b2018 =>> 0xc0000ae008
	c := &b // b  0xc0000b2018 =>> 0xc0000ae008
	fmt.Printf("a值:%d &a-pts:%p type:%T &a-type:%T \n", a, &a, a, &a)
	fmt.Printf("b值:%d b-pts:%p &b-pts:%p type:%T \n", *b, b, &b, b)
	fmt.Printf("c值:%d c-pts:%p *c-pts:%p type:%T \n", **c, c, *c, c)

}
