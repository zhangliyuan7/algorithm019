package Repeat_01

import (
	"math"
)

//220
// abs (nums[i]-nums[j])<=t
// abs (i-j) <=k
// 桶排序
func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}
//good
func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}


func containsNearbyAlmostDuplicateII(nums []int, k int, t int) bool {
	for i := range nums{
		for j:=i+1;j<=i+k && j <len(nums);j++{
			if math.Abs(float64(nums[i]-nums[j])) <= float64(t){
				return true
			}
		}
	}
	return false
}

