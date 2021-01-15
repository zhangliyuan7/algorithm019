package Repeat_01

// 977
// 精华 ！ 厉害
// 有负数的时候 双端最大是一定的，所以双端比较
func sortedSquares(a []int) []int {
	n := len(a)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := a[i]*a[i], a[j]*a[j]; v > w {
			ans[pos] = v
			i++
		} else {
			ans[pos] = w
			j--
		}
	}
	return ans
}

func SortedSquares(nums []int) []int {
	left:=0
	right:=len(nums)-1
	for i:=0;i<len(nums);i++{
		if nums[i]<0{
			left++
		}
		nums[i]=nums[i]*nums[i]
	}
	for left>0{
		left--
		for nums[right]>nums[0]{
			right--
		}
		t:=nums[0]
		tmp:=append([]int{},nums[right+1:]...)
		nums=nums[1:right+1]
		nums=append(nums,t)
		nums=append(nums,tmp...)
	}
	return nums
}
