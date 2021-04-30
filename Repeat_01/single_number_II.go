package Repeat_01





//137 https://leetcode-cn.com/problems/single-number-ii/

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		//32位的total
		total := int32(0)
		for _, num := range nums {
			// 获取每一位的1/0 加到total上 ，每个数字的每一位都加入到total的二进制位上
			total += int32(num) >> i & 1
		}

		// 恰好只有三次 所以和的余数一定是剩余的那个出现一次的数字，用位数复现出来
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

// way II
// 原理 [a,a,a,b,b,b,c,c,c]与[a,a,a,b,b,b,c]差了两个c
// 所以 只要将第一个去重乘以3 即3*（a+b+c）-3*(a+a+a+b+b+b+c)=2c
// 把原数组去重 乘以三 再减去当前数组 就是2c
// go 中去重比较麻烦
func singleNumberII(nums []int) int {
	set:=make(map[int]int)
	right:=0 //被减得数字
	for i:=range nums{
		set[nums[i]]=1
		right+=nums[i]
	}
	sum:=0
	for i:=range set{
		sum+=i
	}
	sum=sum*3
	return (sum-right)/2
}
//II 化简
func singleNumberIII(nums []int) int {
	set:=make(map[int]int)
	right:=0 //被减得数字
	sum:=0
	for i:=range nums{
		if _,ok:=set[nums[i]];!ok{
			sum+=nums[i]
		}
		set[nums[i]]=1
		right+=nums[i]
	}
	return (sum*3-right)/2
}