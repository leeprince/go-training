package test

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