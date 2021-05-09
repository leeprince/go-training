package test

/*
在升序数组中，在原数组中删除重复元素，并返回新的数组长度(不考虑超出新数组后面的元素)
 */
func uniq(arr []int) int  {
	l := len(arr)
	
	s := 1
	for f := 1; f < l ; f++  {
		if arr[f] != arr[f-1] {
			arr[s] = arr[f]
			s++
		}
	}
	return s
}