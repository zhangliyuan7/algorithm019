package Repeat_01

// 852
func peakIndexInMountainArray(arr []int) int {
	a:=0
	for i:=1;i<len(arr);i++{
		if arr[i]<arr[a]{
			break
		}
		a=i
	}
	return a
}

//方法二较为繁琐，利用另一个数组生成bool数组(递增为true ，递减为false) ，返回true false 变化处即可