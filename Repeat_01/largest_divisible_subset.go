package Repeat_01

import (
	"sort"
)

//368
func largestDivisibleSubset(nums []int) (ans []int ){
	sort.Ints(nums)
	ln := len(nums)
	if ln==0{
		return
	}
	dp := make(map[int][]int)
	dp[0]=[]int{nums[0]}
	ans=dp[0]
	for i := 1; i < ln; i++ {
		dp[i]=append(dp[i],nums[i])
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(dp[j]) >= len(dp[i]) {
				copy(dp[i], dp[j])
				dp[i] = append(dp[i], nums[i])
			}
		}
		if len(dp[i])>len(ans){
			ans=dp[i]
		}
	}
	return
}

func largestDivisibleSubsetI(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return []int{}
	}
	sort.Ints(nums)
	memo := make([][]int, l)
	memo[0] = []int{nums[0]}
	result := memo[0]
	for right := 1; right < l; right++ {
		memo[right] = append(memo[right], nums[right])
		for left := 0; left < right; left++ {
			// 存在整除子集中的最大元素的数字
			if nums[right] % nums[left] == 0 && len(memo[right]) <= len(memo[left]){
				copy(memo[right], memo[left]) // 不能用=赋值！
				memo[right] = append(memo[right], nums[right])
			}
		}
		if len(result) < len(memo[right]) {
			result = memo[right]
		}
	}
	return result
}

// faster
func largestDivisibleSubsetII(nums []int)(ans []int) {
	l := len(nums)
	if l == 0 {
		return []int{}
	}
	sort.Ints(nums)
	dp:=make([]int,l)
	for i:=range dp{
		dp[i]=1
	}
	maxValue:=1
	maxLength:=1
	for i:=1;i<l;i++{
		for j:=0;j<i;j++{
			if nums[i]%nums[j]==0&&dp[j]>=dp[i]{
				dp[i]=dp[j]+1
			}
		}
		if dp[i]>maxLength{
			maxLength=dp[i]
			maxValue=nums[i]
		}
	}
	if maxLength == 1 {
		return []int{nums[0]}
	}
	for i := l - 1; i >= 0 && maxLength > 0; i-- {
		if dp[i] == maxLength && maxValue%nums[i] == 0 {
			ans = append(ans, nums[i])
			maxValue = nums[i]
			maxLength--
		}
	}
	return
}
