package Week_08

//493
func reversePairs(nums []int) int {
	if len(nums)<=1{
		return 0
	}
	//var count int
	n:=len(nums)
	middle:=n/2
	left:=append([]int{},nums[:middle]...)
	right:=append([]int{},nums[middle:]...)
	count:=reversePairs(left) + reversePairs(right)
	l:=0
	r:=0
	for l<len(left)&&r<len(right){
		if left[l] > 2*right[r] {
			count+=len(left)-l
			r++
		}else{
			l++
		}
	}
	p1,p2:=0,0
	for i:=0;i<n;i++{
		if p1<len(left)&&( p2 ==len(right)||left[p1]<right[p2]) {
			nums[i]=left[p1]
			p1++
			continue
		}
		nums[i]=right[p2]
		p2++
	}
	return count
}


func reversePairs2(nums []int) int {
	// 利用归并排序的方法计算翻转对的个数
	var mergeCount func(arr []int) int
	mergeCount = func(arr []int) int {
		// 如果arr的长度为0 1 的话， 不构成翻转对的条件， 返回0
		if len(arr) <= 1 {
			return 0
		}
		n := len(arr)
		// 将arr 分成左右两部分， 分别求解
		left := append([]int{}, arr[:n/2]...)
		right := append([]int{}, arr[n/2:]...)
		lLen, rLen := len(left), len(right)         // 左半部分和右半部分的长度
		cnt := mergeCount(left) + mergeCount(right) // 返回翻转对的长度， 并且left 和 right 已经是升序排序

		// 计算翻转对的个数
		l, r := 0, 0
		for l < lLen && r < rLen {
			if left[l] > 2*right[r] {
				cnt += lLen - l // 如果 left[l] > 2*right[r]， 说明left[l+1:]剩余的元素也满足此条件
				r++
			} else {
				l++
			}
		}
		// 原地对arr进行排序
		p1, p2 := 0, 0
		for i := 0; i < n; i++ {
			if p1 < lLen && (p2 == rLen || left[p1] <= right[p2]) {
				arr[i] = left[p1]
				p1++
			} else {
				arr[i] = right[p2]
				p2++
			}
		}
		return cnt
	}
	return mergeCount(nums)
}

