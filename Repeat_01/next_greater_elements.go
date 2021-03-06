package Repeat_01

// 503

func nextGreaterElements(nums []int)[]int {
	//单调栈，存索引
	var s []int
	n:=len(nums)
	r:=make([]int,n)
	for i:=range r{
		r[i]=-1
	}
	//循环数组 两次才能获取到所有元素之后
	for i:=0;i<n*2-1;i++{
		//%n 保证索引正常
		//如果栈为空，则把当前元素放入栈内；
		//如果栈不为空，则需要判断当前元素和栈顶元素的大小：
		//如果当前元素比栈顶元素大：说明当前元素是前面一些元素的「下一个更大元素」，则逐个弹出栈顶元素，直到当前元素比栈顶元素小为止。
		//如果当前元素比栈顶元素小：说明当前元素的「下一个更大元素」与栈顶元素相同，则把当前元素入栈。
		for len(s)!=0&&nums[i%n]>nums[s[len(s)-1]%n]{
			r[s[len(s)-1]%n]=nums[i%n]
			s=s[:len(s)-1]
		}
		s=append(s,i)
	}
	return r
}
