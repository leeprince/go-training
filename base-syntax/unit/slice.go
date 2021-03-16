package unit

import "fmt"

func Slice()  {
	data := []int{0, 1, 2, 3, 4, 10: 8}
	fmt.Println(data)
	s := data[:2:4]
	fmt.Printf("s:%v;s-ptr:%p;data-ptr:%p \n", s, &s[0], &data[0])
	s = append(s, 100, 200)
	fmt.Printf("s:%v;s-ptr:%p;data-ptr:%p \n", s, &s[0], &data[0])
	s = data[:4:5]
	fmt.Printf("s:%v;s-ptr:%p;data-ptr:%p \n", s, &s[0], &data[0])
	s = append(s, 101, 201)
	fmt.Printf("s:%v;s-ptr:%p;data-ptr:%p \n", s, &s[0], &data[0])
}
