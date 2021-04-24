package unit

import "fmt"

type sliceData struct {
	s []int
}

func Slice()  {
	data := []int{0, 1, 2, 3, 4, 10: 8}
	fmt.Println(data)
	s := data[:2:4]
	fmt.Printf("s:%v;s:%d;s-ptr:%p;data-ptr:%p \n", s, s[1], &s[0], &data[0])
	s = append(s, 100, 200)
	fmt.Printf("s:%v;s-ptr:%p;data-ptr:%p \n", s, &s[0], &data[0])
	s = data[:4:5]
	fmt.Printf("s:%v;s-ptr:%p;data-ptr:%p \n", s, &s[0], &data[0])
	s = append(s, 101, 201)
	fmt.Printf("s:%v;s-ptr:%p;data-ptr:%p \n", s, &s[0], &data[0])
	
	sliceChange(data)
	fmt.Println("修改slice值后的结果：", data)
	
	sliceStruct := sliceData{[]int{1, 2, 3}}
	sliceStructChange(sliceStruct)
	fmt.Println("修改slice值后的结果：", sliceStruct)
	
	sliceStructQuote := sliceData{[]int{1, 2, 3}}
	sliceStructChangeQuote(&sliceStructQuote)
	fmt.Println("修改slice值后的结果：", sliceStruct)
}

func sliceChange(data []int)  {
	data[0] = 100
}

func sliceStructChange(data sliceData)  {
	s := data.s
	fmt.Println("s:", s)
	s1 := data.s
	s1[0] = 100
	fmt.Println("s:", s1)
	
	data.s[0] = 101
}

func sliceStructChangeQuote(data *sliceData)  {
	s := data.s
	fmt.Println("s:", s)
	s1 := data.s
	s1[0] = 100
	fmt.Println("s:", s1)
	
	data.s[0] = 101
}
