package test

import (
	"fmt"
	"testing"
)

func TestUniq(t *testing.T) {
	// 简单用例测试
	/*i := []int{1, 1, 2, 2, 3, 4, 5}
	fmt.Println(uniq(i), i)*/
	
	// 多用例测试
	type args struct {
		arr []int
	}
	tests := []struct {
		// 用例名称
		name string
		// 用例参数
		args args
		// 预期的值
		want int
	}{
		{
			name: "在升序数组中，在原数组中删除重复元素，并返回新的数组长度(不考虑超出新数组后面的元素)",
			args: args{
				arr: []int{1, 1, 2, 2, 3, 4, 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniq(tt.args.arr)
			if got != tt.want {
				t.Errorf("uniq() = %v, want %v", got, tt.want)
			}
			fmt.Println(got, tt.args.arr)
		})
	}
}
