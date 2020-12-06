package Week_06

import (
	"sort"
)
//office
func fourSum(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}


func FourSum(nums []int, target int) [][]int {
	var r [][]int
	sort.Ints(nums)
	var delrepeat = make(map[[4]int]bool)
	for i:=0;i<len(nums);i++{
		l:=len(nums)-1
		for k:=i+1;k<l;k++{
			for j:=k+1;j<l;j++{
				for l:=len(nums)-1;l>j;l--{
					sum:=nums[i]+nums[j]+nums[k]+nums[l]
					if sum==target{
						will:=[4]int{nums[i],nums[k],nums[j],nums[l]}
						if _,ok:=delrepeat[will];ok{
							continue
						}
						delrepeat[will]=true
						r=append(r,[]int{nums[i],nums[k],nums[j],nums[l]})
					}
				}
			}
		}
	}
	return r
}
